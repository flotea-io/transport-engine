<template>
	<el-popover
	placement="top"
	:title="carrierName"
	width="350"
	trigger="hover">
	<div slot>
		<div>
			{{eth.toFixed(4)}} ETH
		</div>
		<div>
			{{flt.toFixed(3)}} FLT
		</div>
		{{userAddress}}
	</div>

	<div slot="reference" class="metamask" @click="enableMetamask">
		<img v-if="!enabledMetamask" src="/metamask-logo-92-a-73-d-44-kopia@3x.png" />
		<i v-if="enabledMetamask" class="el-icon-user-solid"></i>
		<span>{{enabledMetamask? "Połączony":"Odblokuj Metamask"}}</span>
		<span v-if="enabledMetamask">{{userType==0? "pasażer": "przewoźnik"}}</span>
	</div>
</el-popover>
</template>


<script>
import Web3 from 'web3';

const blockchainSharedData = {
	enabledMetamask: false,
	userAddress: null,
	carrierName: "",
	userType: 0,
	flt: 0,
	eth: 0,
	isMetamask: true,
	contractsLoaded: false,
	waitForResponse: false,
};

var addresses = {
	wssUrl: "wss://kovan.infura.io/ws/v3/eabd1b929937436ea9a0529c3eae21c8",
	FloteaToken: "0x1C94D0D2F4ED3E8Cd38B0cBA9C8d05e6648fdaa2",
	VotingCarrier: "0xa71eE8c6fb6079FCB8b13Fa81Bb554950809A74C",
	Carriers: "0xC9C5c6Ac5836AaCEA50Fe7fcf286071032cE5451",
	FloteaICO: "0xDa8370898A95446875DaC7AeAd92E5fd5851174D",
	Transport: "0x0D2C2582F68f792E4f7c6b8045D5a9Ae70806526",
};

//["0x444f875c1754Eb221032f8023DC03de2bC2098fF","0x568a7332Ce1236cF0A2433006eAa078d9c924cf7","0x26c68459Ad0ab1298067B020DFEf8FB0EaeBC0Ce"]
//["0x4163636f756e7431000000000000000000000000000000000000000000000000","0x4163636f756e7432000000000000000000000000000000000000000000000000","0x4163636f756e7433000000000000000000000000000000000000000000000000"]
var blockChain = {
	initializedEvents: false,
	contracts: {},
	contractsCall: {},
	mainInit: () =>{
		blockChain.getAccounts();
		blockChain.init();
		blockChain.loadContract("FloteaToken", ()=> blockChain.loadContract("Carriers", ()=>{
			blockchainSharedData.contractsLoaded = true;
			if(blockchainSharedData.enabledMetamask){
				blockChain.getAccountInfo();
			}
		}));
	},
	init: () => {

		if(blockchainSharedData.waitForResponse) return;
		
		if (window.ethereum) {
			blockchainSharedData.waitForResponse = true;
			window.ethereum.autoRefreshOnNetworkChange = false;
			window.web3 = new Web3(window.ethereum);

			const main = async () => {
				await window.ethereum.enable();
			};
			main().catch(err => {
				blockchainSharedData.waitForResponse = false;
				$.jGrowl("Please enable MetaMask, or Opera Wallet.", { position:"center", header: 'Error', life: 10000 });
			});
		}
		else {
			blockchainSharedData.isMetamask = false;
			$.jGrowl("Non-Ethereum browser detected. You should consider trying MetaMask or Opera browser!", { position:"center", header: 'Error', life: 10000 });
		}
	},
	getAccountInfo(){
		blockChain.callContract("Carriers", "getCarrierData", (response)=>{
			//console.log(response, blockchainSharedData.userAddress);
			if(response.exist){
				blockchainSharedData.carrierName = blockChain.fromHex(response.company);
				blockchainSharedData.userType = 1;
			} else {
				blockchainSharedData.carrierName = "";
				blockchainSharedData.userType = 0;
			} 
		}, null, blockchainSharedData.userAddress);
		blockChain.callContract("FloteaToken", "balanceOf", (balance)=>{
			blockchainSharedData.flt = balance/1000;
		}, null, blockchainSharedData.userAddress);
		web3.eth.getBalance(blockchainSharedData.userAddress).then((amount)=>{
			blockchainSharedData.eth = amount/10**18;
		});
	},

	getAccounts(){
		var accountInterval = setInterval(()=>{
			if (web3.currentProvider.selectedAddress !== blockchainSharedData.userAddress) {
				blockchainSharedData.userAddress = web3.currentProvider.selectedAddress;
				if(blockchainSharedData.userAddress != null){
					blockchainSharedData.enabledMetamask = true;
					if(blockchainSharedData.contractsLoaded){
						blockChain.getAccountInfo();
					}
				} else{
					blockchainSharedData.enabledMetamask = false;
				}
			}
		}, 500);
	},

	initWebsocketProvider: () => {
		window.web3wss = new Web3(new Web3.providers.WebsocketProvider(addresses.wssUrl));
	},

	callContract: (contract, method, success, error = null, ...params) => {
		blockChain.contractsCall[contract].methods[method](...params).call().then((result) => {
			if(typeof success == "function")
				success(result);
		}).catch((err) => {
			//console.log(err);
			if(typeof error == "function")
				error(err);
			blockChain.handleErrorMessages(err);
		});
	},

	sendContract(contract, method, sendParams, success, error = null, ...params){
		blockChain.contracts[contract].methods[method](...params).send(sendParams)
		.on('transactionHash', function(hash){
			if(typeof success == "function")
				success(hash);
			//console.log(hash);
		})
		.on('receipt', function(receipt){
			//console.log(receipt);
		})
		.on('error', (err) => {
			if(typeof error == "function")
				error(err);
			blockChain.handleErrorMessages(err);
		});
	},

	sendTransaction(to, value, gas, success, error = null){
		web3.eth.sendTransaction({
			from: blockchainSharedData.userAddress,
			to: to,
			gas: gas,
			value: value
		}).on('transactionHash', function(hash){
			if(typeof success == "function")
				success(hash);
			//console.log(hash);
		})
		.on('receipt', function(receipt){
			//console.log(receipt);
		})
		.on('error', (err) => {
			if(typeof error == "function")
				error(err);
			this.handleErrorMessages(err);
		});
	},

	loadContract(name, fn){
		if(typeof blockChain.contracts[name] != "undefined"){
			if(typeof fn == "function")
				fn(blockChain.contracts[name]);
			return;
		}
		$.getJSON("/"+ name +".json", (e) => {
			if (typeof window.web3 == "object")
				blockChain.contracts[name] = new window.web3.eth.Contract(e.abi, addresses[name]);

			blockChain.contractsCall[name] = new window.web3wss.eth.Contract(e.abi, addresses[name]);
			if(typeof fn == "function"){
				fn(blockChain.contractsCall[name]);

			}
		});
	},
	loadContractAtAddress(name, address, fn){
		if(typeof blockChain.contracts[name+address] != "undefined"){
			if(typeof fn == "function")
				fn(blockChain.contracts[name+address]);
			return;
		}
		$.getJSON("/"+ name +".json", (e) => {
			if (typeof window.web3 == "object")
				blockChain.contracts[name+address] = new window.web3.eth.Contract(e.abi, address);

			blockChain.contractsCall[name+address] = new window.web3wss.eth.Contract(e.abi, address);
			if(typeof fn == "function"){
				fn(blockChain.contractsCall[name+address]);
			}
		});
	},

	fromHex(hex){
		return window.web3wss.utils.hexToUtf8(hex);
	},

	toHex(str){
		let hex = window.web3wss.utils.utf8ToHex(str.toString());
		return hex + '0'.repeat(66-hex.length);
	},

	numberToHex(number){
		return window.web3wss.utils.numberToHex(number);
	},

	handleErrorMessages(error){
		$.jGrowl(error.stack, { position:"center", header: error.message, life: 10000 });
	}
};
export { blockChain };
export default {
	mounted() {
		require("jgrowl");
		blockChain.initWebsocketProvider();
		blockChain.mainInit();			
		this.$emit('getObject', blockChain);
		//console.log(this, blockChain);
	},
	data: function(){
		return blockchainSharedData;
	},
	computed: {
		metamaskState(){
			return {
				enabledMetamask: blockchainSharedData.enabledMetamask,
				userAddress: blockchainSharedData.userAddress,
				carrierName: blockchainSharedData.carrierName,
				userType: blockchainSharedData.userType,
				flt: blockchainSharedData.flt,
				eth: blockchainSharedData.eth,
				isMetamask: blockchainSharedData.isMetamask,
			};
		}
	},
	watch:{
		metamaskState: function(newState){
			this.$emit('metamaskUpdated', newState);	
		}
	},
	methods:{
		enableMetamask: function(){ 
			if(!this.isMetamask){
				let link = "https://metamask.io/";
				if(navigator.userAgent.indexOf("Opera") != -1 || navigator.userAgent.indexOf('OPR') != -1) 
					link = "https://addons.opera.com/en/extensions/details/metamask/";
					else if(navigator.userAgent.indexOf("Chrome") != -1 ) // Chrome and Brave have same addon
						link = "https://chrome.google.com/webstore/detail/nkbihfbeogaeaoehlefnkodbefgpgknn";
					else if(navigator.userAgent.indexOf("Firefox") != -1 ) 
						link = "https://addons.mozilla.org/en-US/firefox/addon/ether-metamask/";
					window.open(link, '_blank');
				}
				else if(!this.enabledMetamask){
					blockChain.init();
				}
			}
		},
	}
	</script>

	<style scoped>
	.metamask{
		border-radius: 20px;
		box-shadow: 0.5px 0.9px 5px 0 rgba(0, 0, 0, 0.1);
		border: solid 1px #ebebeb;
		font-size: 14px;
		font-weight: normal;
		font-style: normal;
		font-stretch: normal;
		line-height: 1.71;
		letter-spacing: normal;
		text-align: center;
		color: #333333;
		cursor: pointer;
		padding: 11px 14px 8px 14px;
		display: inline-block;
	}
	.metamask img{
		width: 18px;
		margin: -5px 8px 0px 3px;
	}

	</style>