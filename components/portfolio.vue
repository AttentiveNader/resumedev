<template>
	<div id="rightpanel">
		<div id="pull-right-btn" href="#" @click="pullRight()" class="pull-btn pull-right-btn"><fa icon="angle-up" />
     	</div>
		<div class="wrapper">
			<div class="box"> 
			   <div class="table">
          <div class="table-cell">
  					<p>worked on the front end of two connected websites with another main web developer(it was his idea) and I built alot of the design, doing so I had to deal with bootstrap/mdbootstrap and of course HTML/CSS & javascript </p>
  					<p><a target="_new" href="http://advancedentity.com">Advanced Entity</a></p>
  					<p><a target="_new" href="http://riftgen.com">Riftgen</a></p>
          </div>
				</div>
			</div>
			<div class="box">
				<div class="table">
          <div class="table-cell">
  					<p>built a url shortener with golang and mysql</p>
  					<a target="_new" href="http://gosho.herokuapp.com"><img src="~/static/scaledlogo.png"></a>
          </div>
				</div>
			</div>
			<div class="box">
        <div class="table">
  				<div class="table-cell">
  					<p><a target="_new" href="http://freecodecamp.org/naderjs">FCC profile</a></p>
  					<p>finished the freecodecamp front end program and made some small projects for it, and this tic tac toe is one of them I developed it to be undefeatable</p>
  					<p><a target="_new" href="https://naderjs.lib.id/ticvue@dev/#/">Tic tac toe</a></p>
  					<!--<p><a target="_new" href="https://lifegamecon.herokuapp.com/">Game of life</a></p>-->
          </div>
				</div>
			</div>
			<div class="box">
				<div class="table">
          <div class="table-cell">
  					<p>you can find more about me on</p>
  					<p><a target="_new" href="https://drive.google.com/open?id=12_6Wo3znBqPQYRYAa1tpB1a5Pq8pm8IT">My Resume PDF</a></p>
  					<p>or</p>
  					<a  target="_new" href="https://github.com/NaderJS"><img class="icon" src="~/static/github.png"/></a>
  					<a  target="_new" href="https://www.linkedin.com/in/nader-atef-6806b2139/"><img class="icon" src="~/static/linkedin.png"/></a>
  					<a  target="_new" href="https://www.facebook.com/profile.php?id=100008930282377"><img class="icon" src="~/static/facebook.png"/></a>
          </div>
				</div>
			</div>
		</div>
	</div>
</template>
<script>
	export default { 
    data(){
      return {
        pulledStatus: false,
        rightPanel: null
      }
    },
    methods:{
      pullRight(){
        var self = this
        if (!self.pulledStatus){
          self.rightPanel.classList.add("expanded")
          self.pulledStatus = !self.pulledStatus
          return
        }
        self.pulledStatus = !self.pulledStatus
        self.rightPanel.classList.remove("expanded") 
      }
    },
    mounted(){
    	 this.rightPanel = document.getElementById('rightpanel')
    }
  }

</script>
<style type="text/css" scoped>
#rightpanel{
	position: fixed;
	left:-100%;
	top: 0px;
	width: 100%;
	height: 100%;
	z-index: 3;
	-webkit-transition: 0.5s; 
    -o-transition: 0.5s; 
    transition: 0.5s;
    font-weight: bolder;
}
.wrapper{	
    padding:20px;
    display: flow-root;
    width: 100%;
    height: 100%; 
    overflow-y: auto; 
}
.pull-right-btn{
	position: absolute;
	top: calc(50% - 20px);
	right:-20px;
	-webkit-animation :2s scaleright infinite ease-in-out ;
	animation :2s scaleright infinite ease-in-out ;
}
@keyframes scaleright{
    0% {
       -webkit-transform: rotate(90deg) scale(1)  ;
    -o-transform: rotate(90deg) scale(1) ;
    -ms-transform: rotate(90deg) scale(1) ;
    -moz-transform: rotate(90deg) scale(1) ; 
            transform:rotate(90deg) scale(1) ;
    }
    50%{

       -webkit-transform: rotate(90deg) scale(1.6) ;
    -o-transform: rotate(90deg) scale(1.6);
    -ms-transform: rotate(90deg) scale(1.6);
    -moz-transform: rotate(90deg) scale(1.6); 
            transform:rotate(90deg) scale(1.6);
    }
    100% {
       -webkit-transform: rotate(90deg) scale(1) ;
    -o-transform: rotate(90deg) scale(1) ;
    -ms-transform: rotate(90deg) scale(1) ;
    -moz-transform: rotate(90deg) scale(1) ; 
            transform:rotate(90deg) scale(1);
    }
  }
  @keyframes scaleleft{
    0% {
       -webkit-transform: rotate(270deg) scale(1)  ;
    -o-transform: rotate(270deg) scale(1) ;
    -ms-transform: rotate(270deg) scale(1) ;
    -moz-transform: rotate(270deg) scale(1) ; 
            transform:rotate(270deg) scale(1) ;
    }
    50%{

       -webkit-transform: rotate(270deg) scale(1.6) ;
    -o-transform: rotate(270deg) scale(1.6);
    -ms-transform: rotate(270deg) scale(1.6);
    -moz-transform: rotate(270deg) scale(1.6); 
            transform:rotate(270deg) scale(1.6);
    }
    100% {
       -webkit-transform: rotate(270deg) scale(1)  ;
    -o-transform: rotate(270deg) scale(1) ;
    -ms-transform: rotate(270deg) scale(1) ;
    -moz-transform: rotate(270deg) scale(1) ; 
            transform:rotate(270deg) scale(1);
    }
  }
 
  .expanded .pull-right-btn{
    -webkit-animation: 2s scaleleft infinite ease-in-out !important ;
         animation: 2s scaleleft infinite ease-in-out !important ;
  }
.icon{
	max-width: 60px !important;
	margin-right: 10px;
}
.expanded{
	position: fixed;
	left:0 !important;
    background:#00daff73;
}
.wrapper .box  div{
  padding-top: 2%;
	width: 100%;
}
.table{
  display: table;
  width: 100%;
  height: 100%;

}
.table-cell {
  display: table-cell; 
  vertical-align: middle;
}
.wrapper .box img{
	max-width: 60%;
}

.wrapper .box{
	float:left;
	width: 48%;
	display: block;
	padding:20px;
	background-color: rgba(255,255,255,0.9);
	backface-visibility: 0.3;
	min-height: 48%;
	margin:5px;
	text-align: center;
}
@media screen and (max-width: 600px){
	.wrapper .box{
		width: 100%;
		padding:5px;
		margin:0 auto;
		margin-bottom: 10px;
	}
}
</style>