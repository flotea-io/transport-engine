<template>
  <div class="d-ib">
    <el-popover trigger="click" placement="top" width="300" ref="monthFrom" v-model="visible[0]">
      <el-row class="month-popover">
        <el-col :span="6" v-for="(month, index) in months" v-bind:data="month" v-bind:key="index">
          <el-button :disabled="monthTo && (monthTo -1 < index)" size="mini" @click="selectMonth('monthFrom', index+1, 0)" type="text">{{month}}</el-button>
        </el-col>
      </el-row>
    </el-popover>

    <el-popover trigger="click" placement="top" width="300" ref="dayFrom" v-model="visible[1]">
      <el-row class="month-popover">
        <el-col :span="3" v-for="(day, index) in getDays('dayFrom')" v-bind:data="day" v-bind:key="index">
          <el-button :disabled="monthFrom == monthTo && dayTo && (dayTo -1 < index)" size="mini" @click="selectDay('dayFrom', index+1)" type="text">{{day}}</el-button>
        </el-col>
      </el-row>
    </el-popover>

    <el-popover trigger="click" placement="top" width="300" ref="dayTo" v-model="visible[3]">
      <el-row class="month-popover">
        <el-col :span="3" v-for="(day, index) in getDays('dayTo')" v-bind:data="day" v-bind:key="index">
          <el-button :disabled="monthFrom == monthTo && (dayFrom -1 > index)" size="mini" @click="selectDay('dayTo', index+1)" type="text">{{day}}</el-button>
        </el-col>
      </el-row>
    </el-popover>

    <el-popover trigger="click" placement="top" width="300" ref="monthTo" v-model="visible[2]">
      <el-row class="month-popover">
        <el-col :span="6" v-for="(month, index) in months" v-bind:data="month" v-bind:key="index">
          <el-button :disabled="(monthFrom -1 > index)" size="mini" @click="selectMonth('monthTo', index+1, 2)" type="text">{{month}}</el-button>
        </el-col>
      </el-row>
    </el-popover>

    <el-button-group >
      <el-button size="mini" v-popover:monthFrom>{{ getMonthName(monthFrom) }}</el-button>
      <el-button :disabled="monthFrom==null" size="mini" v-popover:dayFrom>{{ dayFrom? dayFrom : "Select day" }}</el-button>
      <el-button size="mini">To</el-button>
      <el-button :disabled="monthFrom==null" size="mini" v-popover:monthTo>{{ getMonthName(monthTo) }}</el-button>
      <el-button :disabled="monthTo==null" size="mini" v-popover:dayTo>{{ dayTo? dayTo : "Select day" }}</el-button>
      <el-button size="mini" @click="clear" type="primary" icon="el-icon-circle-close"></el-button>
    </el-button-group>
  </div>
</template>


<script>
  import moment from 'moment';
  export default {
    props: ["data", "open"],
    data() {
      return {
        months: [],
        visible: [false],
        monthFrom: null,
        monthTo: null,
        dayFrom: null,
        dayTo: null 
      };
    },
    methods: {
      selectMonth(month, index, m) {
        this.visible[m] = false;
        this[month] = index;

        this["day"+month.substr(5)] = null;
        this.change();
      },
      selectDay(day, index){
        this[day] = index;
        this.change();
      },
      getMonthName(index){
        if(index == null)
          return "Select month";
        return moment("2000-"+ (index > 9? index : "0"+index) +"-01").format('MMMM');
      },
      getDays(type){
        let month = this["month"+type.substr(3)];
        if(!month) return 0;
        let days = moment("2000-"+ (month > 9 ? month : "0" + month) , "YYYY-MM").daysInMonth();
        return days;
      },
      clear(change = true){
        this.monthFrom = null;
        this.monthTo = null;
        this.dayFrom = null;
        this.dayTo = null;
        if(change)
          this.change();
      },
      prepareData(){
         this.clear(false);
//console.log(this.data);
         return;
         for (var i = 0; i <= 1; i++) {
           if (this.data.length > 3) {
            if(this.data[i].length > 0)
              this["month"+(i==0? "From" : "To")] = this.data[i][0];
            if(this.data[i].length > 1)
              this["day"+(i==0? "From" : "To")] = this.data[i][1];
          }
         }
      },
      change(){
        let ranges = [];
        if(this.dayFrom)
          ranges.push([this.monthFrom, this.dayFrom]);
        else if(this.monthFrom)
          ranges.push([this.monthFrom]);
        if(this.dayTo)
          ranges.push([this.monthTo, this.dayTo]);
        else if(this.monthTo)
          ranges.push([this.monthTo]);
        this.$listeners['update:data'](ranges);
        this.$listeners['change']();
      }
    },
    created: function() { 
      let january = moment("2000-12-01");
      for (var i = 0; i < 12; i++) {
        this.months.push(january.add(1, "month").format('MMMM'));
      }
      this.prepareData();
    }
  }
</script>

<style scoped>
.d-ib{
  display: inline-block;
}
.month-popover{
  margin-top: -14px;
}
.month-popover .el-col{
  margin-top: 14px;  
}
.month-popover button{
  width: 100%;
  text-align: center;
}
</style>