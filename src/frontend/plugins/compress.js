var compress = {

  toHexArray(scheduleObject) {

    let expItems = [
      '{"or":\\[',//0
      '{"and":\\[',//1
      '{"not":\\[',//2
      '"wd":\\[',//3
      '"bhr":\\[',//4
      '"hr":\\[',//5
      '"bev":\\[',//6
      '"eev":\\[',//7
      '"bdate":\\[',//8
      '"edate":\\[',//9
      '\\{',//A
      '(\\]\\})+',//B
      '\\d+',
      '(\\]\\})+,',
      '\\],',
      ','
    ];

    let jsonString = JSON.stringify(scheduleObject);
    // convert numbers to base 13
    jsonString = jsonString.replace(new RegExp(expItems[12], "g"), (char) => parseInt(char).toString(13));
    // ]}+,
    jsonString = jsonString.replace(new RegExp(expItems[13], "g"), function(char) { // D
      let add = 0;
      if (expItems[13].indexOf("+") != -1) { // multiple
        let s = (new RegExp("\\(([^)]+)")).exec(expItems[13])[1];
        s = s.replace(/\\/g, "");
        add = Math.floor(char.length / s.length)
      }
      if (add == 1)
        return "d"; // only once
      else
        return "d" + (add).toString(13) + "d"; // multiply
    });
    // ],
    jsonString = jsonString.replace(new RegExp(expItems[14], "g"), "e");
    // commands 0-A
    for (let i = 0; i <= 10; i++) {
      jsonString = jsonString.replace(new RegExp(expItems[i], "g"), (i).toString(13));
    }
    // ,
    jsonString = jsonString.replace(new RegExp(expItems[15], "g"), "f");
    // remove ending close ]}+
    jsonString = jsonString.replace(new RegExp(expItems[11], "g"), "");

    let returnArray = jsonString.match(/.{1,64}/g); // divide string to 32 bytes

    for (var i = 0; i < returnArray.length; i++) {
      if (i == returnArray.length - 1)
        returnArray[i] = "0x" + returnArray[i] + '0'.repeat(64 - returnArray[i].length); // fill zero to 32 bytes
      else
        returnArray[i] = "0x" + returnArray[i];
    }
    return returnArray;
  },

  fromHexArray(array) {
    let commands = [
      '{"or":[',
      '{"and":[',
      '{"not":[',
      '"wd":[',
      '"bhr":[',
      '"hr":[',
      '"bev":[',
      '"eev":[',
      '"bdate":[',
      '"edate":[',
      '{',
      null, null, null, null,
      ','
    ]
    let str = array.join("").replace(/0x/g, "");
    if (str[str.length - 1] == "0") {
      for (var i = str.length - 1; i >= 0; i--) {
        if (str[i] != "0") {
          str = str.substr(0, i + 1);
          break;
        }
      }
    }
    str += "K" 
    //console.log(str);
    let reverseStr = str.match(/.{1,1}/g).reverse().join("");
    let pom, rec = "",
      state = "start",
      strLength = str.length,
      par = 0;
    let re, lastCommandPosition, returnNumbers, lastNumber = 0;
    for (var i = 0; i < strLength; i++) {
      if (lastNumber != 0) {
        i = strLength - 1;
      }
      let command = parseInt(str[i], 16);
      //console.log(command);
      switch (true) {
        case (command <= 2):
          rec += commands[command];
          par++;
          break;
        case (command <= 9):
          rec += commands[command];
          lastCommandPosition = i;
          [i, re, lastNumber] = this.findNumbers(i, str);
          //console.log(re);
          rec += re;
          break;
        case (command == 10 || command == 15):
          rec += commands[command];
          break;
        case (command == 13):
          pom = reverseStr.indexOf("f", reverseStr.length - i);
          if (pom == -1) {
            pom = lastCommandPosition + 1;
          } else {
            pom = Math.max(strLength - pom, lastCommandPosition + 1);
          }
          rec += parseInt(str.substr(pom, i - pom), 13);
          i++;
          let f = str.indexOf("f", i);
          let d = str.indexOf("d", i);
          if (d == -1) {
            d = i + 3;
          }
          let e = str.indexOf("e", i);
          let repeat = parseInt(str.substr(i, d - i), 13);
          if ((f < d && f != -1) || (e < d && e != -1) || d - i > 2) {
            repeat = 1;
            i--;
          } else {
            i = d;
          }
          par -= repeat - 1;
          for (var r = 0; r < repeat; r++) {
            rec += "]}";
          }
          rec += ",";
          break;
        case (command == 14):
          pom = reverseStr.indexOf("f", reverseStr.length - i);
          if (pom == -1) {
            pom = lastCommandPosition + 1;
          } else {
            pom = Math.max(strLength - pom, lastCommandPosition + 1);
          }
          rec += parseInt(str.substr(pom, i - pom), 13) + "],";
          break;
      }
      if (str[i] == "K") {
        rec += parseInt(str.substr(lastNumber, i), 13);
        for (var p = 0; p <= par; p++) {
          rec += "]}";
        }
      }
    }
    return rec;
  },

  findNumbers(i, str) {
    let numberPosition = -1,
      jsonString = "",
      lastNumberPosition = 0,
      end = 3;
    i++;
    do {
      numberPosition = str.indexOf("f", i);
      let d = str.indexOf("d", i);
      let e = str.indexOf("e", i);
      if (d != -1) numberPosition = Math.min(numberPosition == -1 ? d : numberPosition, d);
      if (e != -1) numberPosition = Math.min(numberPosition == -1 ? e : numberPosition, e);
      if (numberPosition != -1) {
        if (str[numberPosition] == "f") {
          jsonString += parseInt(str.substr(i, numberPosition - i), 13) + ",";
          i = numberPosition + 1;
        } else {
          i = numberPosition - 1;
        }
      } else {
        lastNumberPosition = i;
      }
      if (end-- < 0) break;
    } while (numberPosition != -1 && str[numberPosition] == "f");
    return [i, jsonString, lastNumberPosition];
  }

}

export default compress;