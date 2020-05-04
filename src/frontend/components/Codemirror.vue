<template>
  <client-only placeholder="Codemirror Loading...">
    <codemirror v-model="codeText" 
    :options="cmOption"
    @ready="onCmReady"
    @blur="onCmBlur">
  </codemirror>
</client-only>
</template>

<script>
export default {
  props: ["code", "visible"],
  data() {
    return {
      codeText: "",
      codemirror: null,
      cmOption: {
        styleActiveLine: true,
        lineNumbers: true,
        height: 50,
        line: true,
        keyMap: "sublime",
        matchBrackets: true,
        mode: "application/ld+json",
        lineWrapping: true,
        theme: 'eclipse',
        foldGutter: true,
        gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter"]
      }
    }
  },
  watch:{
    visible: function(){
      this.codemirror.setValue(this._props.code);
    },
    codeText: function(){
      //console.log(this.codemirror.doc.height);
    },
    code: function(){
      this.codeText = this._props.code;
    }
  },
  methods: {
    IsJsonString(str) {
      try {
        var json = JSON.parse(str);
        return (typeof json === 'object');
      } catch (e) {
        return false;
      }
    },
    onCmReady(codemirror) {
      this.codeText = this._props.code;
      this.codemirror = codemirror;
    },
    onCmBlur(codemirror) {
      if(this.IsJsonString(this.codeText)){
        this.$emit("change", JSON.parse(this.codeText));
      }

    }
  }
}
</script>

<style>
.CodeMirror{
  height: auto;
}
</style>