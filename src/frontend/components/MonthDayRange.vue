<template>
	<el-container>
		<el-aside width="140" class="left">
			<el-button type="success" icon="el-icon-plus" size="small" circle v-on:click="addMdRange"></el-button>
			Repeat month day range 
		</el-aside> 
		<el-aside width="auto" class="right">
			<div v-for="(item, index) in monthDayRanges" v-bind:data="item" v-bind:key="index">
				<el-dropdown trigger="click">
					<el-button type="info" size="mini">
						<span v-if="item.not">NOT</span>
						<i class="el-icon-edit"></i>
					</el-button>
					<el-dropdown-menu slot="dropdown" class="fullWidth">
						<el-dropdown-item>
							<el-button :icon="item.not? 'el-icon-circle-check':'el-icon-circle-close'" size="small" type="text" v-on:click="item.not=!item.not">
								{{item.not? "granted":"NOT"}}
							</el-button>
						</el-dropdown-item>
						<el-dropdown-item>
							<el-button icon="el-icon-delete" size="small" type="text" v-on:click="removeDateRange(index)">Remove</el-button>
						</el-dropdown-item>
					</el-dropdown-menu>
				</el-dropdown>
				<month-day-input :data.sync="monthDayRanges[index].data" v-on:change="change" />
				
			</div>
		</el-aside>
	</el-container>
</template>

<script>
	import moment from 'moment';
	import MonthDayInput from '~/components/MonthDayInput.vue';
	export default {
		props: ["data", "open"],
		components:{
			MonthDayInput
		},
		data: function(){
			return {
				monthDayRanges: []
			};
		},
		watch:{
			open: function(before, after){
				this.prepareData();
			},
			monthDayRanges: function() { this.change(); }
		},
		methods: {
			addMdRange() {
				this.monthDayRanges.push({data:[], not: false});
			},
			removeDateRange(index){
				this.monthDayRanges.splice(index, 1);
			},
			change(){
				this.$listeners['update:data'](this.monthDayRanges);
			},
			prepareData(){
				this.monthDayRanges = this.data;
			}
		},
		created: function() { this.prepareData(); }
	}

</script>

<style scoped>
	.left{
		width: 215px;
		margin-top: 5px;
		text-align: left;
	}
	.right{
		text-align: left;
		border-left: 1px solid #EBEEF5;
    	padding-left: 20px;
	}
	.fullWidth li{
    	padding: 0px;
  	}
  	
  	.fullWidth button.el-button{
    	width: 100px;
        text-align: left;
    	padding-left: 12px;
  	}
</style>