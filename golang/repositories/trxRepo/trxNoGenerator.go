package trxrepo

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type autonumberValue struct {
	Prefix      string `json:"prefix"`
	Datatype    string `json:"datatype"`
	SeqValue    int    `json:"seqvalue"`
	LeadingZero int    `json:"leadingzero"`
}

func (ctx trxRepository) GenerateNo(datatype string, prefix string, leadingZero ...int) (code string, err error) {
	var autonumber autonumberValue
	zeroPadding := 0

	if len(leadingZero) > 0 {
		zeroPadding = leadingZero[0]
	}
	fmt.Println("zeroPadding", zeroPadding)
	query := `select data_type, seqvalue from no_generators where data_type= $1`
	//cek data ada ga?
	err = ctx.repo.Db.QueryRow(query, datatype).Scan(&autonumber.Datatype, &autonumber.SeqValue)
	if err != nil {
		fmt.Println(err)

		//klo ga ada insert
		query := `insert into no_generators  (data_type,leadingzero,seqvalue)values($1,$2,$3) returning data_type, leadingzero,coalesce(prefix,'') prefix, seqvalue`
		err := ctx.repo.Db.QueryRow(query, datatype, zeroPadding, 1).Scan(&autonumber.Datatype, &autonumber.LeadingZero, &autonumber.Prefix, &autonumber.SeqValue)
		if err != nil {
			fmt.Println(err)
			log.Println("ERROR INSERT")
			return "", err
		}
	} else {
		autonumber.SeqValue = autonumber.SeqValue + 1
		autonumber.LeadingZero = zeroPadding
		// update
		query = `update no_generators set data_type=$1,leadingzero=$2,seqvalue=$3 where data_type=$4`
		errr := ctx.repo.Db.QueryRow(query, autonumber.Datatype, autonumber.LeadingZero, autonumber.SeqValue, autonumber.Datatype)
		if errr.Err() != nil {
			log.Println("ERROR UPDATE", errr)
			return "", err
		}
	}

	autonumberNo := ""
	if zeroPadding != 0 {
		iSeq, _ := strconv.ParseInt(strconv.Itoa(autonumber.SeqValue), 10, 64)
		lpad := padLeft(iSeq, autonumber.LeadingZero)
		autonumberNo = fmt.Sprintf("%s%s", prefix, lpad)
	} else {
		autonumberNo = fmt.Sprintf("%s%s", prefix, strconv.Itoa(autonumber.SeqValue))
	}
	return datatype + "-" + autonumberNo, nil
}
func padLeft(v int64, length int) string {
	abs := math.Abs(float64(v))
	var padding int
	if v != 0 {
		min := math.Pow10(length - 1)

		if min-abs > 0 {
			l := math.Log10(abs)
			if l == float64(int64(l)) {
				l++
			}
			padding = length - int(math.Ceil(l))
		}
	} else {
		padding = length - 1
	}
	builder := strings.Builder{}
	if v < 0 {
		length = length + 1
	}
	builder.Grow(length * 4)
	if v < 0 {
		builder.WriteRune('-')
	}
	for i := 0; i < padding; i++ {
		builder.WriteRune('0')
	}
	builder.WriteString(strconv.FormatInt(int64(abs), 10))
	return builder.String()
}
