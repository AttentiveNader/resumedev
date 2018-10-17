<template>
    <div id="top-panel">
      <div id="pull-down-btn" href="#" @click="pullDown()" class="pull-btn pull-down-btn rotate"><fa icon="angle-up" />
      </div>
     
      <form v-on:submit.prevent="sendMessage()" class="contact-form">
    	 	<div>
    			<h2>Contact Me</h2>
    		</div>
    	  <div>
    	    <label for="con-name">Full name</label>
    	    <input type="text" name="con-name" id="con-name" v-model="contactForm.name"/>
    	  </div>
    	  <div>
    	    <label for="con-email">Email</label>
    	    <input type="email" name="con-email" id="con-email" v-model="contactForm.address"/>
    	  </div>
    	  <textarea placeholder="Your Message" v-model="contactForm.message"></textarea>
        <div id="response">
          <p>{{message}}</p>
        </div>
    	  <button  @click="sendMessage()" class="send-btn">Send</button>
    		  <p>My Email address is nader_atef80@outlook.com</p>
      </form>
    </div>  
</template>
<script>
  import axios from "axios"
  export default { 
    data(){
      return {
        pulledStatus: false,
        topPanel: null,
        container : null,
        Loading : null,
        response : null,
        contactForm: {
          "name": "",
          "address": "",
          "message": ""
        },
        message: ""
      }
    },
    methods:{
      pullDown(){
        var self = this
        if (!self.pulledStatus){
          self.topPanel.classList.add("expanded")
          self.container.style.display = "block"
          self.container.style.opacity = "1"
          self.pulledStatus = !self.pulledStatus
          return
        }
        self.pulledStatus = !self.pulledStatus
        self.container.style.display = "none"
        self.container.style.opacity = "0"
        self.topPanel.classList.remove("expanded") 
      },
      sendMessage(){
        let self  =  this
        if (this.contactForm.message.length < 1 || this.contactForm.address.length < 1 || this.contactForm.name.length < 1) {      
          this.message = "please fill all  fields"
          response.classList.add("red-resp")
          return
        }
        self.Loading.style.display = "block"
        axios.post("https://na-resume-api.herokuapp.com/send",this.contactForm)
        .then(function(response){
          self.message = response.data.message
          console.log(response)
          self.Loading.style.display = "none"
          self.response.style.display = "block"

        })
        .catch(err => {
          self.message = err
          console.log(err)
          self.Loading.style.display = "none"
        })
      }
    },
    mounted(){
    	 this.topPanel = document.getElementById('top-panel')
       this.container = document.getElementById("container")
       this.Loading = document.getElementById("loading")
       this.response = document.getElementById("response")
    }
  }
</script>
<style type="text/css" scoped>
#top-panel{
    position: fixed;
    top: -100%;
    width: 100%;
    z-index: 6;
    left: 0px;
    height: 100%; 
    -webkit-transition: 0.5s; 
    -o-transition: 0.5s; 
    transition: 0.5s;
    background:#c53f55;
  }
  #response{
    display: none;
    background: lightblue;
    opacity : 0.7;
    text-align: center !important;
    color:black;
    border-radius: 12px;
    padding-top: 5px;
  }
  .red-resp{
    background: red;
    display: block !important;
  }
  .light-resp {
    background:lightblue;
    display: block !important; 
  }
  @media screen and (min-width: 1000px){
    #top-panel{
      position: fixed;
      top: -100%;
      left:10%;
      border-bottom-left-radius: 20px ; 
      border-bottom-right-radius: 20px ; 
      max-width: 80%;
      height: 100%; 
      -webkit-transition: 0.5s; 
      -o-transition: 0.5s; 
      transition: 0.5s;
      background:#c53f55;
    }
  }
   @media screen and (max-height: 600px){
    .contact-form {
      margin-top: 3% !important;
    }
   }
 
  .pull-down-btn{
    margin:0 auto;
    -webkit-transition: 0.4s;
    -o-transition: 0.4s;
    transition: 0.4s;
    position: absolute;
    bottom:-20px;
    left:calc(50% - 20px);
    -webkit-animation: 2s scaledown infinite ease-in-out ;
         animation: 2s scaledown infinite ease-in-out ;
  }
 @-webkit-keyframes scaledown{
    0% {
       -webkit-transform: rotate(180deg) scale(1)  ;
    -ms-transform: rotate(180deg) scale(1) ; 
            transform:rotate(180deg) scale(1) ;
    }
    50%{

       -webkit-transform: rotate(180deg) scale(1.6) ;
    -ms-transform: rotate(180deg) scale(1.6); 
            transform:rotate(180deg) scale(1.6);
    }
    100% {
       -webkit-transform: rotate(180deg)  ;
    -ms-transform: rotate(180deg) ; 
            transform:rotate(180deg) scale(1);
    }
  }
 @keyframes scaledown{
    0% {
       -webkit-transform: rotate(180deg) scale(1)  ;
    -ms-transform: rotate(180deg) scale(1) ; 
            transform:rotate(180deg) scale(1) ;
    }
    50%{

       -webkit-transform: rotate(180deg) scale(1.6) ;
    -ms-transform: rotate(180deg) scale(1.6); 
            transform:rotate(180deg) scale(1.6);
    }
    100% {
       -webkit-transform: rotate(180deg)  ;
    -ms-transform: rotate(180deg) ; 
            transform:rotate(180deg) scale(1);
    }
  }
  @-webkit-keyframes scaleup{
    0% {
       -webkit-transform: rotate(0deg) scale(1)  ;
    -ms-transform: rotate(0deg) scale(1) ; 
            transform:rotate(0deg) scale(1) ;
    }
    50%{

       -webkit-transform: rotate(0deg) scale(1.6) ;
    -ms-transform: rotate(0deg) scale(1.6); 
            transform:rotate(0deg) scale(1.6);
    }
    100% {
       -webkit-transform: rotate(0deg)  ;
    -ms-transform: rotate(0deg) ; 
            transform:rotate(0deg) scale(1);
    }
  }
  @keyframes scaleup{
    0% {
       -webkit-transform: rotate(0deg) scale(1)  ;
    -ms-transform: rotate(0deg) scale(1) ; 
            transform:rotate(0deg) scale(1) ;
    }
    50%{

       -webkit-transform: rotate(0deg) scale(1.6) ;
    -ms-transform: rotate(0deg) scale(1.6); 
            transform:rotate(0deg) scale(1.6);
    }
    100% {
       -webkit-transform: rotate(0deg)  ;
    -ms-transform: rotate(0deg) ; 
            transform:rotate(0deg) scale(1);
    }
  }
 
  .expanded .pull-down-btn{
    -webkit-animation: 2s scaleup infinite ease-in-out !important ;
         animation: 2s scaleup infinite ease-in-out !important ;
  }
  .expanded{
    top:0 !important;
  }
  .contact-form{
    text-align: center;
    max-width: 350px;
    margin: 0 auto;
    margin-top: 15%;
    background: aliceblue;
    padding: 10px;
    border-radius: 12px;
    font-size: 18px;
  }
  .contact-form div{
    margin: 0 auto;
    padding-bottom: 10px;
    margin-bottom: 10px;
    text-align: left;
  }
  .contact-form div label{
    display: block;
    float:left;
    width: 100px;
    padding-top: 5px;
  }
  .send-btn{
    width: 100%;
    margin-bottom: 10px;
  }
  textarea{
    width: 100%;
    height:100px;
    resize: none;
    margin-bottom: 10px;
  }
  input[type=text],input[type=email], textarea{
    font-size: 16px;
    padding : 5px;
    outline: none;
    border:2px solid #ccc;
  }
  input:focus,textarea:focus{
    border:2px solid #d3d3d3 ;
    -webkit-box-shadow: 0 0 5px rgba(81, 203, 238, 1);
            box-shadow: 0 0 5px rgba(81, 203, 238, 1);
  }
</style>