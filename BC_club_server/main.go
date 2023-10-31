package main

import (
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

var mainPageHTML []byte
var loginHTML []byte
var registerHTML []byte
var competitionHTML []byte
var bioResHTML []byte
var bioSugHTML []byte
var cheResHTML []byte
var cheSugHTML []byte
var APPageHTML []byte
var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("mysql", "root:Ton@8177919@tcp(101.43.188.94:3306)/BC_club?charset=utf8")
	if err != nil {
		fmt.Println("failed to log in the database server")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("login success")
	err = db.Ping()
	if err != nil {
		fmt.Println("failed to ping")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ping success")

	defer func(db *sql.DB) {
		_ = os.WriteFile("./running_info.log", []byte("progress were killed at: "+time.Now().String()), os.ModePerm)
		_ = db.Close()
	}(db)

	mainPageHTML, err = os.ReadFile("./web_page/main_page.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	loginHTML, err = os.ReadFile("./web_page/login.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	registerHTML, err = os.ReadFile("./web_page/register.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	competitionHTML, err = os.ReadFile("./web_page/competition.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bioResHTML, err = os.ReadFile("./web_page/bio_res.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bioSugHTML, err = os.ReadFile("./web_page/bio_sug.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cheResHTML, err = os.ReadFile("./web_page/che_res.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cheSugHTML, err = os.ReadFile("./web_page/che_sug.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	APPageHTML, err = os.ReadFile("./web_page/AP_page.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web_page/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web_page/css"))))
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./web_page/res"))))
	http.Handle("/bootstrap-5.1.0/", http.StripPrefix("/bootstrap-5.1.0/", http.FileServer(http.Dir("./web_page/bootstrap-5.1.0"))))
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/AP", APPageHandler)
	http.HandleFunc("/AP/bio/res", bioResHandler)
	http.HandleFunc("/AP/bio/sug", bioSugHandler)
	http.HandleFunc("/AP/che/res", cheResHandler)
	http.HandleFunc("/AP/che/sug", cheSugHandler)
	http.HandleFunc("/competition", competitionHandler)
	http.HandleFunc("/res_sug", resSugHandler)

	socket, err := net.Listen("tcp", ":32347")
	go AcceptConnection(socket)

	err = http.ListenAndServe(":32346", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func AcceptConnection(listener net.Listener) {
	for true {
		connection, err := listener.Accept()
		if err != nil {
			return
		}
		go ReceiveData(connection)
	}
}

func ReceiveData(conn net.Conn) {

	for true {
		method := make([]byte, 1)
		_, err := conn.Read(method)
		if err != nil {
			return
		}

		receiveSignal := make([]byte, 1)
		receiveSignal[0] = 0

		switch method[0] {
		case 0:
			target := make([]byte, 1)
			_, err := conn.Read(target)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			lengthBytes := make([]byte, 4)
			_, err = conn.Read(lengthBytes)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			length := binary.LittleEndian.Uint32(lengthBytes)
			htmlBytes := make([]byte, length)
			_, err = conn.Read(htmlBytes)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			err = conn.Close()
			if err != nil {
				fmt.Println(err.Error())
			}
			switch target[0] {
			case 0: /* BiologyResource = 0, BiologySuggestion = 1, ChemistryResource = 2, ChemistrySuggestion = 3 */
				err = os.WriteFile("./web_page/bio_res_extra.html", htmlBytes, os.ModePerm)
				if err != nil {
					fmt.Println(err.Error())
					receiveSignal[0] = 1
				}
				break
			case 1:
				err = os.WriteFile("./web_page/bio_sug_extra.html", htmlBytes, os.ModePerm)
				if err != nil {
					fmt.Println(err.Error())
					receiveSignal[0] = 1
				}
				break
			case 2:
				err = os.WriteFile("./web_page/che_res_extra.html", htmlBytes, os.ModePerm)
				if err != nil {
					fmt.Println(err.Error())
					receiveSignal[0] = 1
				}
				break
			case 3:
				err = os.WriteFile("./web_page/che_sug_extra.html", htmlBytes, os.ModePerm)
				if err != nil {
					fmt.Println(err.Error())
					receiveSignal[0] = 1
				}
				break
			}
			break
		case 1:
			nameLengthBytes := make([]byte, 4)
			_, err = conn.Read(nameLengthBytes)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			nameLength := binary.LittleEndian.Uint32(nameLengthBytes)
			fileNameBytes := make([]byte, nameLength)
			_, err = conn.Read(fileNameBytes)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fileName := string(fileNameBytes)

			dataLengthBytes := make([]byte, 4)
			_, err = conn.Read(dataLengthBytes)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			dataLength := binary.LittleEndian.Uint32(dataLengthBytes)
			fileDataBytes := make([]byte, dataLength)
			_, err = conn.Read(fileDataBytes)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			err = os.WriteFile("./web_page/res/"+fileName, fileDataBytes, os.ModePerm)
			if err != nil {
				fmt.Println(err.Error())
				receiveSignal[0] = 1
			}
			break

		case 2:
			lengthBytes := make([]byte, 4)
			_, err = conn.Read(lengthBytes)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			length := binary.LittleEndian.Uint32(lengthBytes)
			data := make([]byte, length)
			_, err = conn.Read(data)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			err = os.WriteFile("./web_page/competition_info.json", data, os.ModePerm)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			break
		}
		_, err = conn.Write(receiveSignal)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func mainPageHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, err := writer.Write(mainPageHTML)
		if err != nil {
			print(err.Error())
		}
	}
}

func loginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, err := writer.Write(loginHTML)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		var loginInfo userInfoStructure
		err = json.Unmarshal(data, &loginInfo)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Method: " + loginInfo.Method)
		fmt.Println("UserName: " + loginInfo.UserName)
		fmt.Println("Password: " + loginInfo.Password)

		res, err := db.Query("select password from test_user_info where username=?", loginInfo.UserName)
		if err != nil {
			fmt.Println(err.Error())
		}
		if res.Next() {
			var password string
			err = res.Scan(&password)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("database: " + password + " user: " + loginInfo.Password)
			if password == loginInfo.Password {
				_, err = writer.Write([]byte("RIGHT"))
				if err != nil {
					fmt.Println(err.Error())
				}
			} else {
				_, err = writer.Write([]byte("WRONG"))
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		} else {
			_, err = writer.Write([]byte("NOT_EXIST"))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

type userInfoStructure struct {
	Method   string `json:"method"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func registerHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, err := writer.Write(registerHTML)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		var registerInfo userInfoStructure
		err = json.Unmarshal(data, &registerInfo)
		if err != nil {
			fmt.Println(err.Error())
		}
		res, err := db.Query("select * from test_user_info where username=?", registerInfo.UserName)
		if err != nil {
			fmt.Println(err.Error())
		}
		if res.Next() {
			_, err = writer.Write([]byte("OCCUPIED"))
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}
		_, err = db.Exec("insert into test_user_info(username, password) values(?, ?)", registerInfo.UserName, registerInfo.Password)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		_, err = writer.Write([]byte("OK"))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func competitionHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, err := writer.Write(competitionHTML)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if req.Method == "POST" {
		info, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if string(info) == "GET_COMPETITION" {
			data, err := os.ReadFile("./web_page/competition_info.json")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			_, err = writer.Write(data)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}

func bioResHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, err := writer.Write(bioResHTML)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func bioSugHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, err := writer.Write(bioSugHTML)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func cheResHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, err := writer.Write(cheResHTML)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func cheSugHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, err := writer.Write(cheSugHTML)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func resSugHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		info, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var innerPage []byte
		switch string(info) {
		case "BIO_SUGGESTION":
			innerPage, err = os.ReadFile("./web_page/bio_sug_extra.html")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			break
		case "BIO_RESOURCE":
			innerPage, err = os.ReadFile("./web_page/bio_res_extra.html")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			break
		case "CHE_SUGGESTION":
			innerPage, err = os.ReadFile("./web_page/che_sug_extra.html")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			break
		case "CHE_RESOURCE":
			innerPage, err = os.ReadFile("./web_page/che_res_extra.html")

			if err != nil {
				fmt.Println(err.Error())
				return
			}
			break
		}
		_, err = writer.Write(innerPage)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func APPageHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		_, err := writer.Write(APPageHTML)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
