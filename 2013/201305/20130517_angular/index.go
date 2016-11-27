package main

var Index = `
<!doctype html>
<html lang="zh_CN" ng-app="Calc">
	<head>
    <!-- Angular Material Dependencies -->
    <script src="angular.min.js"></script>
    <script src="angular-animate.min.js"></script>
    <script src="angular-aria.min.js"></script>

    <script src="angular-material.min.js"></script>
          
    <link rel="stylesheet" href="angular-material.min.css">
				
<style>
md-toolbar md-icon.md-default-theme {
  color: white; }
.md-grid-list {
  margin: 8px; }
.gray {
  background: #f5f5f5; }
.green {
  background: #b9f6ca; }
.yellow {
  background: #ffff8d; }
.blue {
  background: #84ffff; }
.purple {
  background: #b388ff; }
.red {
  background: #ff8a80; }
md-grid-tile {
  transition: all 400ms ease-out 50ms; }
</style>

<script type="text/javascript" 
  src="toolbar.js">
</script>


	</head>
	
	<body ng-controller="AppCtrl">

    <div>
        <md-content>
          <md-toolbar>
            <div class="md-toolbar-tools">
              <md-button class="md-icon-button" aria-label="Settings"
                   ng-click="toggleLeft()" class="md-primary">
                <md-icon md-svg-icon="menu.svg"></md-icon>
              </md-button>
              <h2>
                <span>训练：加减法</span>
              </h2>
            </div>
          </md-toolbar>  
        </md-content>     
    </div>
 
 <div>
 <md-grid-list
        md-cols-sm="1" md-cols-md="2" md-cols-gt-md="6"
        md-row-height-gt-md="1:1" md-row-height="2:2"
        md-gutter="12px" md-gutter-gt-sm="8px" >
        
    <md-grid-tile ng-repeat="fa in fat.app track by $index" class="blue"
        md-rowspan="1" md-colspan="1" md-colspan-sm="1">
        <md-input-container flex>
          <label ng-model="fat.app[$index]">[[fat.app[index] ]]</label>
          <input ng-model='[[fa.x]]+"+"+[[fa.y]]+"="' onBlur="return Get_value([$index]);">
        </md-input-container>
        <md-grid-tile-footer>
             <div ng-controller='tm_ctl'>
               <h3>[[clock.now]]</h3>
             </div>
        </md-grid-tile-footer>
    </md-grid-tile>

   
  </md-grid-list>
 </div>     
    
<div layout="column" style="height:500px;">
  <section layout="row" flex>
    <md-sidenav class="md-sidenav-left md-whiteframe-z2" md-component-id="left">

        <md-toolbar class="md-theme-left">
          <md-button class="md-icon-button" aria-label="Settings"
                     ng-click="toggleLeft()" class="md-primary">
             <md-icon md-svg-icon="menu.svg"></md-icon>
          </md-button>
          <h2 class="md-toolbar-tools">设置</h2>
        </md-toolbar>

      <md-content ng-controller="LeftCtrl" layout-padding>
        <form>
          <md-input-container>
            <label for="testInput">Test input</label>
            <input type="text" id="testInput"
                   ng-model="data" md-sidenav-focus>
          </md-input-container>
        </form>
        <md-button ng-click="close()" class="md-primary">
          Close Sidenav Right
        </md-button>
      </md-content>
    </md-sidenav>
  </section>
</div>



	</body>
</html>

`;