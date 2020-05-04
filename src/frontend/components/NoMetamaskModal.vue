<template>
	<el-dialog
	:visible.sync="visible"
	:close-on-click-modal="false"
	:close-on-press-escape="false"
	width="100%" class="text-center metamaskDialog">
	<img src="/metamask-logo-92-a-73-d-44-kopia-4.png">
	<h2>Uruchom portfel Metamask</h2>

	<p>
		{{part1}}
		<el-link type="primary" :href="link" target="_blank">pobierz metamask tutaj</el-link>
		{{part2}}
	</p>

</el-dialog>
</template>
<script>
export default {
	props: ["text", "metamaskDialogVisible"],
	mounted: function(){
		if(navigator.userAgent.indexOf("Opera") != -1 || navigator.userAgent.indexOf('OPR') != -1) 
			this.link = "https://addons.opera.com/en/extensions/details/metamask/";
		else if(navigator.userAgent.indexOf("Chrome") != -1 )
			this.link = "https://chrome.google.com/webstore/detail/nkbihfbeogaeaoehlefnkodbefgpgknn";
		else if(navigator.userAgent.indexOf("Firefox") != -1 ) 
			this.link = "https://addons.mozilla.org/en-US/firefox/addon/ether-metamask/";
		let p = this.text.indexOf("link")
		this.part1 = this.text.substr(0,p);
		this.part2 = this.text.substr(p+4);
	},
	data: function () {
		return {
			visible: false,
			part1: "",
			part2: "",
			link: "https://metamask.io/"
		};
	},
	watch:{
		metamaskDialogVisible: function(n){
			//this.visible = n;
		},
	}
}
</script>