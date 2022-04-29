/*
SQL схема
CREATE TABLE [dbo].[mailqueue](
	[id] [int] NULL,
	[email] [nchar](64) NULL,
	[text] [ntext] NULL,
	[status] [tinyint] NULL,
	[statusText] [ntext] NULL
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

insert into mailqueue (id,email,text) values (1,'tal82@bk.ru','pismo no 1');
insert into mailqueue (id,email,text) values (2,'tal82@bk.ru','pismo no 2');
insert into mailqueue (id,email,text) values (3,'tal82@bk.ru','pismo no 3');
insert into mailqueue (id,email,text) values (4,'tal82@bk.ru','pismo no 4');
insert into mailqueue (id,email,text) values (5,'tal82@bk.ru','pismo no 5');

select *  from dbo.mailqueue 
*/
package main

//используемые модули
import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/alexbrainman/odbc"
	"net"
	"net/smtp"
	"net/mail"
	 "crypto/tls"
)

//переменные
var (
	id int
	email string
	text string	
	mailStatus int
	mailResult string
)

//константы (конфигурация)
const (
	confMailAccount = "tut.pochta@gmail.com"
	confMailPass = "parol.ot.pochti"
	confMailSubject = "Письмо из тестового приложения Golang"
	confMailSerer = "smtp.gmail.com:465"
	
	confSqlUser = "sa"
	confSqlServer = "mssqlcon"
	confSqlPass = "123"
)

//функция отправки письма
func mysendmail(paramTo string,paramMsg string) (int,string) {

	var errmsg string
	
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			fmt.Errorf("panic in getReport %s", err)
		}
	}()

	from := mail.Address{"", confMailAccount}
    to   := mail.Address{"", paramTo}
    subj := confMailSubject
    body := paramMsg

    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subj

    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    servername := confMailSerer

    host, _, _ := net.SplitHostPort(servername)

    auth := smtp.PlainAuth("",confMailAccount,confMailPass, host)

    // TLS config
    tlsconfig := &tls.Config {
        InsecureSkipVerify: true,
        ServerName: host,
    }
  
    conn, err := tls.Dial("tcp", servername, tlsconfig)
    if err != nil {
		errmsg=err.Error()
        log.Panic(err)
		return 0,errmsg
    }

    c, err := smtp.NewClient(conn, host)
    if err != nil {
		errmsg=err.Error()
        log.Panic(err)
		return 0,errmsg
    }

    // Auth
    if err = c.Auth(auth); err != nil {
		errmsg=err.Error()
        log.Panic(err)
		return 0,errmsg
    }

    // To && From
    if err = c.Mail(from.Address); err != nil {
		errmsg=err.Error()
        log.Panic(err)
		return 0,errmsg
    }

    if err = c.Rcpt(to.Address); err != nil {
		errmsg=err.Error()
        log.Panic(err)
		return 0,errmsg
    }

    // Data
    w, err := c.Data()
    if err != nil {
		errmsg=err.Error()
        log.Panic(err)
		return 0,errmsg
    }

    _, err = w.Write([]byte(message))
    if err != nil {
		errmsg=err.Error()
        log.Panic(err)
		return 0,errmsg
    }

    err = w.Close()
    if err != nil {
		errmsg=err.Error()
        log.Panic(err)
		return 0,errmsg
    }

    c.Quit()
	
	return 1,"good"
}

//главный цикл программы
func main() {
	//подключаемся к базе
	connectionString := fmt.Sprintf("DSN=%s;Uid=%s;Pwd=%s",confSqlServer,confSqlUser,confSqlPass)
	db, err := sql.Open("odbc", connectionString)
	if err != nil {
		fmt.Println("Error in connect DB")
		log.Fatal(err)
	}
	
	//запрос писем
	query := "select id,trim(email) email,text  from dbo.mailqueue where status is NULL"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	
	//цикл по очереди писем
	for rows.Next() {
		if err := rows.Scan(&id,&email,&text); err != nil {
			log.Fatal(err)
		}
		log.Println(email,text)	
		
		//отправляем письмо
		mailStatus,mailResult = mysendmail(email,text)	
		//выводим результат отправки
		log.Println(mailStatus,mailResult)		
		
		//запись результата отправки		
		_, err := db.Exec(`update dbo.mailqueue set status=?, statusText=? where id=?;`,mailStatus,mailResult,id)
		if err != nil {
			log.Fatal(err)
		}		
	}
	
	//освобождаем объекты
	defer rows.Close()	
	defer db.Close()	
}
