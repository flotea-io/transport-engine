<template>
  <div class="root">
  <div class="global">
    <nuxt />
  </div>
  <div class="footer">
        <img src="/FE_logo.png">
        <img src="/RP_logo.png">
        <img src="/Slaskie_logo.png">
        <img src="/UE_logo2.png">
    </div>
  </div>
</template>

<style lang="scss">
@import '~/node_modules/bootstrap/scss/_functions.scss';
@import '~/node_modules/bootstrap/scss/_variables.scss';
@import '~/node_modules/bootstrap/scss/_mixins.scss';
@import '~/node_modules/bootstrap/scss/_utilities.scss';

html, .root{
  height: 100%;
}
body{
  height: 100%;
  margin: 0px;
  font-family: "Open Sans";
}
#__nuxt, #__layout{
  height: 100%;
}
.header{
  height: 30px;
}
.global{
  display: flex;
  height: calc(100% - 60px);
  flex-direction: column;
}
.footer{
  height: 50px;
  text-align: center;
  padding-top: 10px;
  img{
        margin: 0px 10px;
  }
}
.el-dialog__body{
  word-break: break-word; 
}
.metamaskDialog{
  max-width: 30em;
  margin: 0px auto;
  p{
    line-height: 24px;
    a.el-link{
      margin-top: -5px;
    }
  }
}
.content{
flex: 1;
}

.info-box{
    border-radius: 3px;
    border: solid 1px #b8daff;
    background-color: #cce5ff;
    padding: 20px 18px 16px 18px;
    max-width: 37em;
    margin: 2em auto;
    color: #004085;
    border-left: 9px solid #004085;
    font-size: 14px;
    text-align: left;
}

h1{
  font-weight: 600;
}

.btn-primary {
  font-family: "Open Sans";
  font-weight: bold;
  font-size: 14px;
  background: none;
  color: #027ebc;
  border-radius: 4px;
  padding: 6px 12px;
  line-height: 1.42857143;
  border: 3px solid #027ebc; }

.btn-primary:focus, .btn-primary:active, .btn-primary:active:hover, .btn-primary:active:focus, .btn-primary:hover {
  background-color: #027ebc;
  border: 3px solid #027ebc;
  background-image: none;
  text-shadow: none;
  box-shadow: none; 
  color: white;
}

@media (min-width: 1200px){
  .container-block {
    margin: 0px auto;
  }
}
</style>