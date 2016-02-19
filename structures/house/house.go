package house

import (
	"encoding/xml"
)

//const dateformat = "2006-01-02"

// Сведения по номерам домов улиц городов и населенных пунктов, номера земельных участков и т.п
type XmlObject struct {
	XMLName    xml.Name `xml:"House" db:"as_house_"`
	POSTALCODE string  `xml:"POSTALCODE,attr,omitempty" db:"postal_code"`
	IFNSFL     int      `xml:"IFNSFL,attr,omitempty" db:"ifns_fl"`
	TERRIFNSFL int      `xml:"TERRIFNSFL,attr,omitempty" db:"terr_ifns_fl"`
	IFNSUL     int      `xml:"IFNSUL,attr,omitempty" db:"ifns_ul"`
	TERRIFNSUL int      `xml:"TERRIFNSUL,attr,omitempty" db:"terr_ifns_ul"`
	OKATO      string  `xml:"OKATO,attr,omitempty" db:"okato"`
	OKTMO      *string  `xml:"OKTMO,attr,omitempty" db:"oktmo"`
	UPDATEDATE string   `xml:"UPDATEDATE,attr" db:"update_date"`
	HOUSENUM   string  `xml:"HOUSENUM,attr,omitempty" db:"house_num"`
	ESTSTATUS  int      `xml:"ESTSTATUS,attr" db:"est_status"`
	BUILDNUM   string  `xml:"BUILDNUM,attr,omitempty" db:"build_num"`
	STRUCNUM   string  `xml:"STRUCNUM,attr,omitempty" db:"struc_num"`
	STRSTATUS  int      `xml:"STRSTATUS,attr" db:"str_status"`
	HOUSEID    string   `xml:"HOUSEID,attr" db:"house_id"`
	HOUSEGUID  string   `xml:"HOUSEGUID,attr" db:"house_guid"`
	AOGUID     string   `xml:"AOGUID,attr" db:"ao_guid"`
	STARTDATE  string   `xml:"STARTDATE,attr" db:"start_date"`
	ENDDATE    string   `xml:"ENDDATE,attr" db:"end_date"`
	STATSTATUS int      `xml:"STATSTATUS,attr" db:"stat_status"`
	NORMDOC    *string  `xml:"NORMDOC,attr,omitempty" db:"norm_doc"`
	COUNTER    int      `xml:"COUNTER,attr" db:"counter"`
}

// схема таблицы в БД

// const tableName = "as_house"
// const elementName = "House"

func Schema(tableName string) string {
	return `CREATE TABLE ` + tableName + ` (
    house_id UUID NOT NULL,
    postal_code VARCHAR(6),
		ifns_fl INT,
		terr_ifns_fl INT,
		ifns_ul INT,
		terr_ifns_ul INT,
		okato VARCHAR(11),
		oktmo VARCHAR(11),
		update_date TIMESTAMP NOT NULL,
		house_num VARCHAR(20),
		est_status INT NOT NULL,
		build_num VARCHAR(20),
		struc_num VARCHAR(20),
		str_status INT,
		house_guid UUID NOT NULL,
		ao_guid UUID NOT NULL,
		start_date TIMESTAMP NOT NULL,
		end_date TIMESTAMP NOT NULL,
		stat_status INT NOT NULL,
		norm_doc UUID,
		counter INT NOT NULL,
		PRIMARY KEY (house_id));`
}

/*
func Export(w *sync.WaitGroup, c chan string, db *sqlx.DB, format *string) {

	defer w.Done()
	helpers.DropAndCreateTable(schema, tableName, db)

	var format2 string
	format2 = *format
	fileName, err2 := helpers.SearchFile(tableName+"_", format2)
	if err2 != nil {
		fmt.Println("Error searching file:", err2)
		return
	}

	pathToFile := format2 + "/" + fileName

	xmlFile, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	total := 0
	var inElement string
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			inElement = se.Name.Local

			if inElement == elementName {
				total++
				var item XmlObject

				// decode a whole chunk of following XML into the
				// variable item which is a ActualStatus (se above)
				err = decoder.DecodeElement(&item, &se)
				if err != nil {
					fmt.Println("Error in decode element:", err)
					return
				}

				query := `INSERT INTO ` + tableName + ` (house_guid,
					postal_code,
					ifns_fl,
					terr_ifns_fl,
					ifns_ul,
					terr_ifns_ul,
					okato,
					oktmo,
					update_date,
					house_num,
					est_status,
					build_num,
					struc_num,
					str_status,
					house_id,
					ao_guid,
					start_date,
					end_date,
					stat_status,
					norm_doc,
					counter
					) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
						$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
						$21)`

				db.MustExec(query,
					item.HOUSEGUID,
					item.POSTALCODE,
					item.IFNSFL,
					item.TERRIFNSFL,
					item.IFNSUL,
					item.TERRIFNSUL,
					item.OKATO,
					item.OKTMO,
					item.UPDATEDATE,
					item.HOUSENUM,
					item.ESTSTATUS,
					item.BUILDNUM,
					item.STRUCNUM,
					item.STRSTATUS,
					item.HOUSEID,
					item.AOGUID,
					item.STARTDATE,
					item.ENDDATE,
					item.STATSTATUS,
					item.NORMDOC,
					item.COUNTER)

				c <- helpers.PrintRowsAffected(elementName, total)
			}
		default:
		}

	}

	//fmt.Printf("Total processed items in AddressObjects: %d \n", total)
}
*/
