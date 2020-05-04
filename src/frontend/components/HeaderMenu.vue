<template>
	<el-menu :default-active="activeIndex" mode="horizontal" text-color="#707070" active-text-color="#333333" @select="handleSelect">
		<el-submenu index="0" class="show-small">
			<i slot="title" class="el-icon-s-unfold el-icon--right"></i>
			<el-menu-item index="1">Znajdź przejazd</el-menu-item>
			<el-menu-item index="2">Twoje bilety</el-menu-item>
			<el-menu-item index="4">Twoje trasy</el-menu-item>
			<el-menu-item index="3">Tworzenie trasy</el-menu-item>
			<el-menu-item index="5">Flotea</el-menu-item>
			<div class="help2">
				<span class="size" @click="zoom">A+</span>
				<span class="contrast" @click="switchCss"><span>A</span><span>A</span></span>
			</div>
		</el-submenu>	
		<el-menu-item class="hide-small" index="1">Znajdź przejazd</el-menu-item>
		<el-menu-item class="hide-small" index="2">Twoje bilety</el-menu-item>
		<el-menu-item class="hide-small" index="4">Twoje trasy</el-menu-item>
		<el-menu-item class="hide-small" index="3">Tworzenie trasy</el-menu-item>
		<Blockchain @metamaskUpdated="(s) => this.$emit('metamaskUpdated', s)" @getObject="(b) => this.$emit('getObject', b)"></Blockchain>
		<img class="eu-head" src="/UE_logo2.png" />
		<el-menu-item class="right hide-small" index="5">Flotea</el-menu-item>
		<div class="help">
			<span class="size" @click="zoom">A+</span>
			<span class="contrast" @click="switchCss"><span>A</span><span>A</span></span>
		</div>
	</el-menu>
</template>

<script>

import Blockchain from "~/components/Blockchain.vue";

export default {
	components: {
		Blockchain
	},
	data () {
		return {
			titles: ["Znajdź przejazd","Twoje bilety","Tworzenie trasy","Twoje trasy"],
			contrast: false,
			files: ['24d78bb4a09868ec66b3.css','82d0f7cf1a9a9d12b9c6.css','093d9c2167484aa59751.css','2039b3064e4935901aa3.css','ab72c26ba82316759ac9.css','aed1aa5e711f5214cb4a.css'],
			zoomSelect: 0,
		}
	},
	head () {
		return {
			title: "Flotea Engine - "+this.titles[parseInt(this.activeIndex)-1],
		}
	},
	created: function() {
		if (process.browser) {
			this.contrast = localStorage.getItem('contrast') == "true";
			this.zoomSelect = parseInt(localStorage.getItem('zoom'));
			if(isNaN(this.zoomSelect))
				this.zoomSelect = 0;
			this.toggleBodyClass(true, "zoom-"+this.zoomSelect);
			var head  = document.getElementsByTagName('head')[0];
			for (var i = 0; i < this.files.length; i++) {
				var link  = document.createElement('link');
				link.id = "contrast"+i;
				link.rel  = 'stylesheet';
				link.type = 'text/css';
				link.disabled = true;
				link.href = '/'+this.files[i];
				head.appendChild(link);
			}
			if(this.contrast)
				this.applyCss(this.contrast);
		}
	},
	props: ["activeIndex"],
	methods:{
		zoom(){
			if (process.browser) {
				this.toggleBodyClass(false, "zoom-"+this.zoomSelect);
				this.zoomSelect = this.zoomSelect > 1? 0 :this.zoomSelect+1;
				this.toggleBodyClass(true, "zoom-"+this.zoomSelect);
				localStorage.setItem('zoom', this.zoomSelect);
			}
		},
		toggleBodyClass(addClass, className) {
			if (process.browser) {
				const el = document.body;
				if (addClass) {
					el.classList.add(className);
				} else {
					el.classList.remove(className);
				}
			}
		},
		switchCss(){
			this.contrast = !this.contrast;
			this.applyCss(this.contrast);
		},
		handleSelect(key, keyPath) {
			let urls = ["/","/passenger/tickets","/carrier/create-route", "/carrier/routes", 
			process.env.NODE_ENV !== 'production'? "http://localhost:8010" : "https://flotea.com"];
			location.href = urls[key-1];
		},
		applyCss(contrast){
			if (process.browser) {
				for (var i = 0; i < this.files.length; i++) {
					document.getElementById("contrast"+i).disabled = !contrast;
				}
				localStorage.setItem('contrast', contrast);
			}
		}
	}
}
</script>

<style lang="scss">
.zoom-0{
	zoom: 1;
}
.zoom-1{
	zoom: 1.2;
}
.zoom-2{
	zoom: 1.4;
}
@media (min-width: 1048px){
	.show-small{
		display: none;
	}
}
@media (max-width: 1048px){
	.hide-small{
		display: none;
	}
}
ul.el-menu{
	>.el-menu-item.is-active{
		background-color: #f5f5f5;
		border-bottom-color: #00bff3!important;
	}
	.right{
		float: right;
	}
	.metamask{
		float: right;
		margin: 8px 8px 0px 8px;
	}
}
.eu-head{
	float: right;
	margin-top: 14px;
}
@media (max-width: 546px){
	.eu-head{
		display: none;
	}
}

.help{
	margin-left: 6px;
	margin-top: 14px;
	display: inline-block;
	width: 88px;
	&:focus {
		outline: none;
	}
	user-select: none;
}
.help2{
	display: none;
}
@media (max-width: 424px){
	.help{
		display: none;
	}
	.help2{
		display: flex;
		margin-left: 9px;
		justify-content: space-around;
		margin-top: 4px;
	}
}
.size{
	cursor: pointer;
}
.contrast{
	display: inline-block;
	outline: 1px solid black;
	span{
		display: inline-block;
		padding: 2px 0px;
		background-color: white;
		color: black;
		cursor: pointer;
		width: 22px;
		text-align: center;
	}
	span+span{
		background-color: black;
		color: white;
	}
}
</style>