<template>
	<div id="rightpanel">
		<div id="pull-right-btn" href="#" @click="pullRight()" class="pull-btn pull-right-btn"><fa icon="angle-up" />
     	</div>
		<div class="wrapper">
			<div class="box"> 
				<div>
					<p>worked on the front end of two related site with another main web developer(it was his idea)</p>
					<p><a target="_new" href="http://advancedentity.com">Advanced Entity</a></p>
					<p><a target="_new" href="http://riftgen.com">Riftgen</a></p>
				</div>
			</div>
			<div class="box">
				<div>
					<p>built a url shortener with golang and mysql</p>
					<a target="_new" href="http://gosho.herokuapp.com"><img src="~/static/scaledlogo.png"></a>
				</div>
			</div>
			<div class="box">
				<div>
					<p><a target="_new" href="http://freecodecamp.org/naderjs">FCC profile</a></p>
					<p>finished the freecodecamp front end program and made some small projects for it, these are 2 of them</p>
					<p><a target="_new" href="https://naderjs.lib.id/ticvue@dev/#/">Tic tac toe</a></p>
					<p><a target="_new" href="https://lifegamecon.herokuapp.com/">Game of life</a></p>
				</div>
			</div>
			<div class="box">
				<div>
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
	-webkit-transform:rotate(90deg);
        -ms-transform:rotate(90deg);
            transform:rotate(90deg);
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
.expanded .pull-right-btn{
	-webkit-transform:rotate(270deg);
        -ms-transform:rotate(270deg);
            transform:rotate(270deg);
}
.wrapper .box  div{
	display: table-cell;
	vertical-align: middle;
	width: 100%;
}
.wrapper .box img{
	max-width: 60%;
}

.wrapper .box{
	float:left;
	width: 48%;
	display: table;
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