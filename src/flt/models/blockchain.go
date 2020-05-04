package models

import (
	//"math/big"
	//"log"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Blockchain struct {
	Id    int    `orm:"auto"`
	Block 	uint64
	Index	uint64
}


func init() {
	orm.RegisterModel(new(Blockchain))
}

func GetLastBlock() (uint64, error){
	o := orm.NewOrm()
	blockchain := Blockchain{Id: 1}

	err := o.Read(&blockchain)

	if err == orm.ErrNoRows {
		fmt.Println("No last block found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		return blockchain.Block, err
	}
	return 0, err
}

func UpdateLastBlock(newBlock uint64) (){
	//return
	o := orm.NewOrm()
	blockchain := Blockchain{Id: 1}	
	blockchain.Block = newBlock
	if num, err := o.Update(&blockchain, "Block"); err == nil {
        _ = num
        //fmt.Println(num)
    }
}
