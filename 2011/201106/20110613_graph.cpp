#include <boost/graph/adjacency_list.hpp>

using namespace boost;

int main(){
  adjacency_list<> g;
  adjacency_list<>::vertex_descriptor v1 = add_vertex(g);
  adjacency_list<>::vertex_descriptor v2 = add_vertex(g);
  adjacency_list<>::vertex_descriptor v3 = add_vertex(g);
  adjacency_list<>::vertex_descriptor v4 = add_vertex(g);
}

