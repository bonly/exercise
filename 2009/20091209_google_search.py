'''
Created on 2011-3-4

@author: bonly
'''
from xgoogle.search import GoogleSearch
  
if __name__ == '__main__':
      gs = GoogleSearch("quick and dirty")
      gs.results_per_page = 50
      results = gs.get_results()
      for res in results:
         print res.title.encode("utf8")
         print res.desc.encode("utf8")
         print res.url.encode("utf8")
         print

