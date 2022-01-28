package main

import (
	"database/sql"
	"encoding/json"
	"example.com/m/v2/ProjectL0/cmd/nats"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

var db *sql.DB
var err error
var JsFromCh DataJs
var i Id
var Cash map[string]DataJs
type Id struct {
	Delivery_id int
	Payment_id  int
	Items_id    int
}

type DataJs struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost uint   `json:"delivery_cost"`
		GoodsTotal   uint   `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []Item `json:"items"`

	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerId        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	Shardkey          string `json:"shardkey"`
	SmId              int    `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OofShard          string `json:"oof_shard"`
}

type Item struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       uint   `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        uint   `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  uint   `json:"total_price"`
	NmId        uint   `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      uint   `json:"status"`
}

func PushIntoCash(Map map[string]DataJs, Data DataJs) {

	var (DeliveryId, PaymentId, ItemsId int)

	rows, err := db.Query("SElECT delivery_id,payment_id,items_id,order_uid,track_number,entry,locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard FROM main")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {

		err := rows.Scan(&DeliveryId,&PaymentId,&ItemsId,&Data.OrderUid, &Data.TrackNumber, &Data.Entry, &Data.Locale, &Data.InternalSignature,&Data.CustomerId,&Data.DeliveryService, &Data.Shardkey, &Data.SmId, &Data.DateCreated, &Data.OofShard)
		if err != nil {
			fmt.Println(err)
			continue
		}

		rows, err = db.Query("SElECT name,phone,zip,city,address,region,email FROM delivery WHERE delivery_id = $1",DeliveryId)
		defer rows.Close()
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			err := rows.Scan(&Data.Delivery.Name, &Data.Delivery.Phone, &Data.Delivery.Zip, &Data.Delivery.City, &Data.Delivery.Address, &Data.Delivery.Region, &Data.Delivery.Email)
			if err != nil {
				fmt.Println(err)
				continue
			}

		}

		rows, err = db.Query("SElECT transaction,request_id,currency,provider,amount,payment_dt,bank,delivery_cost,goods_total,custom_fee FROM payment WHERE payment_id = $1",PaymentId)

		defer rows.Close()
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			err := rows.Scan(&Data.Payment.Transaction, &Data.Payment.RequestId, &Data.Payment.Currency, &Data.Payment.Provider, &Data.Payment.Amount, &Data.Payment.PaymentDt, &Data.Payment.Bank, &Data.Payment.DeliveryCost, &Data.Payment.GoodsTotal, &Data.Payment.CustomFee)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		rows, err = db.Query("SElECT chrt_id,track_number,price,rid,name,sale,size,total_price,nm_id,brand,status FROM items WHERE items_id = $1",ItemsId)

		defer rows.Close()
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			var I Item
				err := rows.Scan(&I.ChrtId, &I.TrackNumber, &I.Price, &I.Rid, &I.Name, &I.Sale, &I.Size, &I.TotalPrice, &I.NmId, &I.Brand, &I.Status)
				if err != nil {
					fmt.Println(err)
					continue
				}

			Data.Items = append(Data.Items,I)
			}

		}

	Map[Data.OrderUid] = Data
	}


func TakeId()  {

	rows, err := db.Query("SELECT delivery_id FROM delivery WHERE zip = $1", JsFromCh.Delivery.Zip)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&i.Delivery_id)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	for _, d := range JsFromCh.Items {
		rows, err := db.Query("SELECT items_id FROM items WHERE chrt_id = $1", d.ChrtId)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&i.Items_id)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	rows, err = db.Query("SELECT payment_id FROM payment WHERE transaction = $1", JsFromCh.Payment.Transaction)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&i.Payment_id)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	defer rows.Close()

}

func PushIntoDB() {

	for _, d := range JsFromCh.Items {
		result, err := db.Exec("CALL insert_data_items($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", d.ChrtId, d.TrackNumber, d.Price, d.Rid, d.Name, d.Sale, d.Size, d.TotalPrice, d.NmId, d.Brand, d.Status)
		if err != nil {
			panic(err)
		}
		PrintUpdate, _ := result.RowsAffected()
		fmt.Println(PrintUpdate)
	}

	result, err := db.Exec("CALL insert_data_payment($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", JsFromCh.Payment.Transaction, JsFromCh.Payment.RequestId, JsFromCh.Payment.Currency, JsFromCh.Payment.Provider, JsFromCh.Payment.Amount, JsFromCh.Payment.PaymentDt, JsFromCh.Payment.Bank, JsFromCh.Payment.DeliveryCost, JsFromCh.Payment.GoodsTotal, JsFromCh.Payment.CustomFee)
	if err != nil {
		panic(err)
	}
	PrintUpdate, _ := result.RowsAffected()
	fmt.Println(PrintUpdate)

	result, err = db.Exec("CALL insert_data_delivery($1,$2,$3,$4,$5,$6,$7)", JsFromCh.Delivery.Name, JsFromCh.Delivery.Phone, JsFromCh.Delivery.Zip, JsFromCh.Delivery.City, JsFromCh.Delivery.Address, JsFromCh.Delivery.Region, JsFromCh.Delivery.Email)
	if err != nil {
		panic(err)
	}
	PrintUpdate, _ = result.RowsAffected()
	fmt.Println(PrintUpdate)

	TakeId()

	result, err = db.Exec("CALL insert_data_main($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)",i.Delivery_id,i.Payment_id,i.Items_id,JsFromCh.OrderUid, JsFromCh.TrackNumber,JsFromCh.Entry,JsFromCh.Locale,JsFromCh.InternalSignature,JsFromCh.CustomerId,JsFromCh.DeliveryService,JsFromCh.Shardkey,JsFromCh.SmId,JsFromCh.DateCreated,JsFromCh.OofShard)
	if err != nil {
		panic(err)
	}
	PrintUpdate, _ = result.RowsAffected()
	fmt.Println(PrintUpdate)
}

func GetData(w http.ResponseWriter, r *http.Request){
	param := mux.Vars(r)
	t,_ := template.ParseFiles("/home/zhora/Desktop/goolang-book/ProjectL0/templates/index.html")
	Cash[param["id"]] = JsFromCh

	err = t.ExecuteTemplate(w, "users", JsFromCh.Delivery)
	if err != nil {
		return
	}

}

func main() {
	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		5432, "localhost", "testuser", "password", "testdb")

	db, err = sql.Open("postgres", pg_con_string)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	defer db.Close()

	Cash = make(map[string]DataJs)
	c := make(chan []byte)

	//При запуске сервиса идет пополнение КЭШа из БД
	PushIntoCash(Cash,JsFromCh)

	fmt.Println("Данные из Мапы",Cash)

	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/users/{id}",GetData).Methods("GET")
		log.Fatal(http.ListenAndServe(":8080",r))
	}()

	go nats.ConnectAndListening(c)

	//Получение данных с сервера, а так же их добавление
	//в КЭШ и БД при условии, что таких данных не было
	for val := range c {

		err := json.Unmarshal(val, &JsFromCh)

		if err != nil {
			fmt.Println(err)
		}

		if _,ok := Cash[JsFromCh.OrderUid];!ok{

			Cash[JsFromCh.OrderUid] = JsFromCh
			PushIntoDB()
		}



	}

}
