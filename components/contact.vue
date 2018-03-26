<template>
    <div id="top-panel">
      <div id="pull-down-btn" href="#" @click="pullDown()" class="pull-btn pull-down-btn"><span> </span> 
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
    	    <input type="email" name="con-email" id="con-email" v-model="contactForm.email"/>
    	  </div>
    	  <textarea placeholder="Your Message" v-model="contactForm.message"></textarea>
        <p id="message">{{message}}</p>
    	  <button  @click="sendMessage()" class="send-btn">Send</button>
    		  <p>My Email address is nader_atef80@outlook.com</p>
      </form>
    </div>  
</template>
<script>
  export default { 
    data(){
      return {
        pulledStatus: false,
        topPanel: null,
        container : null,
        contactForm: {
          name: "",
          email: "",
          message: ""
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
          self.pulledStatus = !self.pulledStatus
          return
        }
        self.pulledStatus = !self.pulledStatus
        self.container.style.display = "none"
        self.topPanel.classList.remove("expanded") 
      },
      sendMessage(){

      }
    },
    mounted(){
    	 this.topPanel = document.getElementById('top-panel')
       this.container = document.getElementById("container")
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
    background:#00daff;
  }

  @media screen and (min-width: 1000px){
  	#top-panel{
  		position: fixed;
 	  	top: -100%;
 	  	left:10%;
	    width: 100%;
	    border-bottom-left-radius: 20px ; 
	    border-bottom-right-radius: 20px ; 
	    max-width: 80%;
	    height: 100%; 
	    -webkit-transition: 0.5s; 
	    -o-transition: 0.5s; 
	    transition: 0.5s;
	    background:#00daf3;
  	}
  }
   @media screen and (max-height: 600px){
    .contact-form {
      margin-top: 4% !important;
    }
   }
  .pull-down-btn{
    margin:0 auto;
    width: 60px;
    height: 60px;
    border-radius: 50% ;
    position: absolute;
    bottom:-30px;
    transform:rotate(90deg);
    left:calc(50% - 30px);
    background: #2c3e50 ;
    transition: 0.5s ;
  }
  .pull-down-btn:hover{
        box-shadow: 0 0 6px black,0 0  6px black ,0 0  6px black ,0 0  6px black ;
  }

  .pull-down-btn:active{
    box-shadow: 0 0 6px black,0 0  6px black ,0 0  6px black ,0 0  6px black ;
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
  	box-shadow: 0 0 5px rgba(81, 203, 238, 1);
  }
</style>