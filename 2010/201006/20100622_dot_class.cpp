/**
 * @file
 */

/**
 * @mainpage
 * 加"\l"结尾表示左对齐
 @dot
 digraph css
 {
   fontname = "文泉驿等宽正黑" fontsize = 16
   Animal [label = "{用户| +name: string\l | +getName():string}" shape="record"]
 }
 @enddot
*/

/** 
 * @file
 * @addtogroup mygr
 * @{
 * @dot
 digraph css
 {
   node [
      fontname = "文泉驿等宽正黑"
      fontsize = 8
      shape = "record"
   ]
   edge [
      fontname = "文泉驿等宽正黑"
      fontsize = 8
   ]
   Myclass [label = "{MyClass | | init()}"]
   main -> Myclass
 }
 * @enddot
 * @}

*/

/**
 * @addtogroup writer
 * @dot
 digraph G {  
   nodesep=0.8;  
   node [ fontname="文泉驿等宽正黑", fontsize=8, shape="record" ]  
   edge [  
     fontsize=8  
     arrowhead="empty"  
     ]  

   Animal [  
     label = "{Animal|+ name: String\l+ age: Integer\l|+ die(): void\l}"  
   ]  

   subgraph clusterAnimalImpl {  
     label="Package animal.impl"  
     Dog [  
       label = "{Dog||+ bark(): void\l}"  
     ]  

     Cat [
       label = "{Cat||}"
     ]
     {rank = same; Dog; Cat}
   }
   Dog -> Animal
   Cat -> Animal

   edge [ arrowhead = "none" headlabel = "0..*" taillabel = "0..*"]
   Dog -> Cat
 }
 @enddot

 rank=same 时,dog/cat时横着放,否则会竖着放
*/

