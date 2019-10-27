import Vue from 'vue'
new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data(){
        return {
            list: ['1','2','3', '24', '50', '60', "80", "90"]
        }
    }
});