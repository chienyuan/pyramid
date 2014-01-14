
#include "Game.h"
#include <iostream>
using namespace std;
#include <fstream>
#include <math.h>

#ifdef MSDOS
#include <conio.h>
#endif

#include <stdlib.h>

// for qsort use compare function
int compare(const void* a,const void* b)
{
	if(*(unsigned *)a>*(unsigned *)b) return 1;
	else return -1;
}

int main()
{
	int		gameover;
	Game	*Gptr;
	char	ans;

	Gptr=new Game();


	cout<<"Welcome to Pyramid game "<<endl;
	cout<<"The Rules of Pyramid: Players alternate take out one to three "<<endl;
	cout<<"tokens in one line. Who takes the last one token is loser"<<endl<<endl;

	do{
	gameover=0;
	Gptr->Reset();
	cout<<"Do you want to move first(y/n)?";
	cin>>ans;
	cin.ignore(256,'\n');
	
	Gptr->Display();


	if(ans=='y'||ans=='Y')
	{
		cout<<"Ok!You move first."<<endl;
		while(!gameover)
		{
			Gptr->Yourmove(gameover);
			if(gameover)
				cout<<" I win"<<endl;
			else{
				Gptr->Mymove(gameover);
				if(gameover) 
					cout<<" You win"<<endl;
			}
			Gptr->FreeNodes();
	
		}
	}else{
		cout<<"Ok!I move first"<<endl;
		while(!gameover)
		{
			Gptr->Mymove(gameover);
			if(gameover)
				cout<<" You win"<<endl;
			else{
				Gptr->Yourmove(gameover);
				if(gameover) cout<<" I win"<<endl;
			}
			Gptr->FreeNodes();
	
		}
	}
	cout<<"Game Over"<<endl;
	cout<<"Do you want to play again(y/n)?";
	cin>>ans;
	cin.ignore(256,'\n');
	}while(ans=='y'||ans=='Y');
		
	delete Gptr;
	


// below is the code for finding win pattern
/*	ofstream win("win.txt"),lose("lose.txt");
	int		gameover;
	unsigned		*Win=new unsigned[65535];
	unsigned		*Loss=new unsigned[65535];
	unsigned		iWin,iLoss;
	
	Game	*Gptr;
	char	ans;

	Gptr=new Game();


	cout<<"Welcome to Pyramid game by prince"<<endl;
	Gptr->FindWin(Win,Loss,iWin,iLoss);
	cout<<"The number of win: "<<iWin;
	cout<<endl;
	cout<<"The number of Loss:"<<iLoss;
	for(unsigned i=0;i<iWin;i++)
	{
		win<<Win[i]<<" ";
		if(!(i%9))
			win<<endl;
	}
	
	for(i=0;i<iLoss;i++)
	{
		lose<<Loss[i]<<" ";
		if(!(i%9))
			lose<<endl;
	}

	win.close();
	lose.close();
	

	ifstream win("win.txt");
	ofstream win2("wins.txt");
	unsigned Win[6300];
	for(int i=0;i<6299;i++)
		win>>Win[i];

	qsort(Win,6299,sizeof(unsigned *),compare);

	for(i=0;i<6299;i++)
	{
		win2<<Win[i]<<" ";
		cout<<Win[i]<<" ";
		if(!(i%9))
		{
			win2<<endl;
			cout<<endl;
		}
	}
*/
}
