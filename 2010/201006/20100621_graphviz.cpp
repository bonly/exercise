/**
 * @file 
 * @brief 注释图的使用
 */

/**
digraph是有向图,graph是无各图
@dot
digraph nodir{
  node [shape=record];
  main -> socket;
  }
@enddot
*/

/*! \mainpage
 *   
 *Class relations expressed via an inline dot graph:
 *\dot
 *       digraph example {
 *           node [shape=record, fontname=Helvetica, fontsize=10];
 *           b [ label="class B" URL="\ref B"];
 *           c [ label="class C" URL="\ref C"];
 *           b -> c [ arrowhead="open", style="dashed" ];
 *       }
 *\enddot
 *      Note that the classes in the above graph are clickable
 *      (in the HTML output).
 **/
