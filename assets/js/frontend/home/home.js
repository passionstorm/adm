import Vue from 'vue'
new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data(){
        return {
            list: ['1','2','3', '24']
        }
    }
});