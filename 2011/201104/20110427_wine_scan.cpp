#include <windows.h>
#include <iostream>
#include <ctime>
#include <string>

class ScanContents{
	public:
		int StartX;
		int StartY;
		int DeductX;
		int DeductY;
		int CompareX;
		int CompareY;
		HDC Hdc;
		
		ScanContents(
		   int startX, int startY, HDC hdc, int compareX = 0,
		   int compareY=0, int deductX=0, int deductY=0){
		   	  StartX = startX;
		   	  StartY = startY;
		   	  CompareX = compareX;
		   	  CompareY = compareY;
		   	  DeductX = deductX;
		   	  DeductY = deductY;
		   	  Hdc = hdc;
		  }		
};

std::string chosenColour;
void MainScan(ScanContents scan);
void CheckColour(COLORREF pixel, int x, int y);
bool ColourMatch(COLORREF pixel);
	
int main(){
	  std::string gameWindow;
	  std::cout << "Enter game window to triggerbot" << std::endl;
	  //std::getline(std::cin, gameWindow);
	  gameWindow = "notepad";//"Call of duty 4";
	
	  //HWND appWnd = FindWindow(0, gameWindow.c_str());
	  HWND appWnd = FindWindow(gameWindow.c_str(), 0);
	  RECT rcClientPositioning;
	  
	  while(!appWnd){
		   //WinExec("clear", SW_SHOW); system("CLS");//
		   appWnd = FindWindow(0, gameWindow.c_str());
		   std::cout << "Looking for " << gameWindow << std::endl;
		   	Sleep(1000);		
	  }
	  std::cout << "Found " << gameWindow << std::endl;
	  
	  while(atoi(chosenColour.c_str()) < 1|| atoi(chosenColour.c_str()) >3 ){
	  	//WinExec("clear", SW_SHOW);system("CLS");//
	  	std::cout << "Choose which color to trigger against\n" <<
	  		"1.Red\n2.Green\n3.Blue"
	  		<< std::endl;
	  		std::getline(std::cin, chosenColour);
	  }
	  
	  std::string color;
	  if (chosenColour == "1") color = "Red";
	  else if (chosenColour == "2") color = "Green";
	  else if (chosenColour == "3") color = "Blue";
	  WinExec("clear", SW_SHOW);//system("CLS");//
	  
	  std::cout << "Triggerbot ONLINE, hover over " << color << " to shoot..." << std::endl;
	  
	  GetWindowRect(appWnd, &rcClientPositioning); 
	  
	  HDC hdcMain = GetDC(HWND_DESKTOP);
	  
	  int startingX = rcClientPositioning.right - ((rcClientPositioning.right - rcClientPositioning.left)/2);
	  int startingY = rcClientPositioning.bottom - ((rcClientPositioning.bottom - rcClientPositioning.top)/2);	  
	  
	  ScanContents scan(startingX, startingY, hdcMain, 30, 40, -30, -30);
	  MainScan(scan);
	  
	  system("PAUSE");
}

void MainScan(ScanContents scan){
	int debugRuntime = clock();
	while(true){
		for (int y = scan.StartY+scan.DeductY; y<scan.StartY+scan.CompareY; y+=3){
		  for (int x = scan.StartX+scan.DeductX; x < scan.StartX+scan.CompareX; x+=3){			
				//Sleep(100);
				//SetCursorPos(x, y);
				
				CheckColour(GetPixel(scan.Hdc, x, y), x, y);
				if(GetAsyncKeyState(VK_DELETE)){
					exit(0);
				}
		 }
	  }
	  std::cout << "Took " << clock() - debugRuntime << " milliseconds to scan whole area" << std::endl;
	  debugRuntime = clock();
	}
	
}


void CheckColour(COLORREF pixel, int x, int y){
	if(ColourMatch(pixel)){
		mouse_event(MOUSEEVENTF_LEFTDOWN, x, y, 0, 0);
		mouse_event(MOUSEEVENTF_LEFTUP, x, y, 0, 0);
	}
}

bool ColourMatch(COLORREF pixel){
	int r=GetRValue(pixel);
	int b=GetBValue(pixel);
	int g=GetGValue(pixel);
	
	//RED
	if (chosenColour == "1"){
		if (r > 50 && g < 40 && b < 40){
			return true;
		}
	}
	//Green
	else if (chosenColour == "2"){
		if (r > 40 && g > 50 && b < 40){
			return true;
		}
	}	
	//BLUE
	else if (chosenColour == "3"){
		if (r < 40 && g < 40 && b > 50){
			return true;
		}
	}	
	return false;
}

//wineg++ file -lgdi32
