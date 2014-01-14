// Game.cpp: implementation of the Game class.
//
//////////////////////////////////////////////////////////////////////

#include "Game.h"
#include <iostream>
using namespace std;
#include <stdlib.h>
#include <math.h>

#ifdef MSDOS
#include <conio.h>
#endif

#include <stdio.h>
#include <fstream>

//////////////////////////////////////////////////////////////////////
// Construction/Destruction
//////////////////////////////////////////////////////////////////////

Game::Game()
{
	int i;
	int temp120[15]={4,8,11,13,14,3,7,10,12,2,6,9,1,5,0};
	int temp240[15]={14,12,9,5,0,13,10,6,1,11,7,2,8,3,4};
	int temp90[15]={0,5,9,12,14,1,6,10,13,2,7,11,3,8,4};
	int temp120r[15]={4,3,2,1,0,8,7,6,5,11,10,9,13,12,14};
	int temp240r[15]={14,13,11,8,4,12,10,7,3,9,6,2,5,1,0};
	s=0;
	for(i=0;i<15;i++){
		rotate120[i]=temp120[i];
		rotate120r[i]=temp120r[i];
		rotate240[i]=temp240[i];
		rotate240r[i]=temp240r[i];
		rotate90[i]=temp90[i];
		power2[i]=(int)pow(2,i);   // assign the access table for power of 2
	}

	Root=new NODE();
	Root->cnum=0;
	Root->endnode=0;
	NodeBuff=new NODE[MAX_SIZE_MEM];
	if(!NodeBuff)
	{
		cerr<<"momory error"<<std::endl;
		exit(-1);
	}
	count=0;
	board=0;
	GenBaseMove();
	ifstream win("wins.txt");
	for(i=0;i<6299;i++)
		win>>Win[i];
	win.close();

}

Game::~Game()
{
	delete Root;
	delete NodeBuff;
}

int Game::AlphaBeta(NODE *p, int a, int b, int depth, int turn,unsigned myBoard)
{
	NODE	*q;
	int		c,temp;

	unsigned	w,i;

	Moving(*p,myBoard);
	if((p->endnode==1)||(depth==0))
	{
		temp=Evaluate(myBoard);
		UnMoving(*p,myBoard);
		if(turn==ME)
			return temp;
		else
			return -temp;
	}

	MoveGen(p,turn,myBoard);
	w=p->children;

	for(i=0;i<p->cnum;i++,w++){
		q=NodeBuff+w;
		c=AlphaBeta(q,a,b,depth-1,-turn,myBoard);
		if(turn==ME){
			if(c==1){
				p->bestchild=w;
				UnMoving(*p,myBoard);
				return 1;
			}

		}
		else{
			if(c==-1){
				p->bestchild=w;
				UnMoving(*p,myBoard);
				return -1;
			}
		}
	}
		
	UnMoving(*p,myBoard);
	return (turn==ME)?a:b;
}

// a binary search just for Evaluate function use 
int BiSearch(unsigned a[],int n,unsigned x)
{
	int first=0;
	int last=n-1;
	int mid = 0;
	int found=0;

	while((first<=last)&&!found)
	{
		mid=(first+last)/2;
		if(x<a[mid])
			last=mid-1;
		else if(x>a[mid])
			first=mid+1;
		else
			found=1;
	}
	if(!found)
		return -1;
	else
		return mid;
}

int Game::Evaluate(unsigned myBoard)
{
	int i,flag=0;
	
	for(i=0;i<15;i++)   // calc how many token in board
	{
		if(myBoard&1)
			flag++;
		myBoard>>=1;

	}
	
	
	if(flag==15||flag==13)
		return 1;
	else if(flag==14)
		return -1;
	
	
	if(BiSearch(Win,6299,myBoard)>0)   // using a winner patter
		return -1;
	else
		return 0 ;
}

// minus the redundance 
int Game::RotateTest(NODE *p,unsigned move,unsigned myBoard)
{
	unsigned r120,r240,r90,r120r,r240r,j;

	r120=Rotate(120,move|myBoard);
	r240=Rotate(240,move|myBoard);
	r120r=Rotate(12090,move|myBoard);
	r240r=Rotate(24090,move|myBoard);
	r90=Rotate(90,move|myBoard);
	
	for(j=0;j<p->cnum;j++){
		if(r120==(NodeBuff[p->children+j].step|myBoard)){
			r++;
			return 1;   }
		if(r240==(NodeBuff[p->children+j].step|myBoard)){
			r++;
			return 1;
		}
		if(r120r==(NodeBuff[p->children+j].step|myBoard)){
			r++;
			return 1;
		}
		if(r240r==(NodeBuff[p->children+j].step|myBoard)){
			r++;
			return 1;
		}
		if(r90==(NodeBuff[p->children+j].step|myBoard)){
			r++;
			return 1;
		}
	}
	return 0;
}			

// generate possible move
void Game::MoveGen(NODE *p, int turn,unsigned myBoard)
{
	int i;
	NODE	*w;

	p->cnum=0;
	s++;
	for(i=0;i<63;i++)
	{
		if(!(Move[i]&myBoard)){  //find valid move

			if(RotateTest(p,Move[i],myBoard))
				continue;

	
			if(count==MAX_SIZE_MEM-1)  // for debug using
			{
				#ifdef DEBUG
					cerr<<"memmory loack"<<std::endl;
					memlack=1;    // memory lack
				#endif
				return;
			}
			(p->cnum)++;
			w=NodeBuff+(++count);  // The child start from 1 not 0
			
			w->step=Move[i];
			w->endnode=IfGameOver(*w,myBoard);
			if(p->cnum==1){
				p->children=count;
				p->bestchild=count;
			}
		}
	}
}

// for generate winner patter use
// not use in the game program
void Game::UnMoveGen(NODE *p, int turn,unsigned myBoard)
{
	int i;
	NODE	*w;

	p->cnum=0;
	s++;
	for(i=0;i<63;i++)
	{
		if(!(Move[i]&myBoard)){  //find valid unmove

			if(RotateTest(p,Move[i],myBoard))
				continue;

	
			if(count==MAX_SIZE_MEM)
			{
#ifdef DEBUG
				cerr<<"error memory over"<<std::endl;
				memlack=1;
#endif
				return;
			}
			(p->cnum)++;
			w=NodeBuff+(++count);  // The child start from 1 not 0
			
			w->step=Move[i];
			w->endnode=IfGameOver(*w,myBoard);
			if(p->cnum==1){
				p->children=count;
				p->bestchild=count;
			}
		}
	}
#ifdef DEBUG
	std::cout<<std::endl<<"board="<<board<<" redundace="<<r<<std::endl;
	getch();
#endif

}

// show the board 
void Game::Display()
{
	int i,j,k=0,m,n,w;
	unsigned temp;

	std::cout<<std::endl;
	m=0;
	w=0;
	
	for(i=4;i>=0;i--){   
		for(k=0;k<i;k++)
				std::cout<<" ";
		n=0;	
		for(j=0;j<5-i;j++){
			m+=n;
			temp=(unsigned)power2[m];
			if(temp&board)		
				std::cout<<"* ";
			else
			{
				if(m<10)
					std::cout<<m<<" ";
				else
					std::cout<<(char)(55+m)<<" ";
			}
			if(j==0)	n=4;
			else		n--;
		}
		
		m=++w;
		std::cout<<std::endl;
	}

	std::cout<<std::endl<<std::endl;
}


int Game::GameOver(unsigned myBoard)
{

	if(myBoard==(unsigned)(pow(2,15)-1))
		return 1;
	else
		return 0;
}

int Game::IfGameOver(NODE w,unsigned myBoard)
{
	int result;
	Moving(w,myBoard);
	result=GameOver(myBoard);
	UnMoving(w,myBoard);
	return result;

}

void Game::Yourmove(int &gameover)
{
	int flag=1,i,j,mx,k;
	unsigned temp;
	char str[80];
	
	if(!gameover)
	{
		while(flag)
		{
			std::cout<<"What is your move?"<<std::endl;
			std::cout<<"X=";
			
			temp=0;
			
			i=0;
			
			cin.getline(str,80);
			
			j=0;
			k=0;
			while(j<3){
				j++;
				while(*(str+k)==' ') //delete the space
					k++;
				if(*(str+k)==0) // if the end of string then break
					break;

				if(*(str+k)>=48 && *(str+k)<=57)   // if input number
					mx=*(str+k)-48;
				else if(*(str+k)>=65 && *(str+k)<=69)  // if input A,B,C..
					mx=*(str+k)-55;
				else if(*(str+k)>=97 && *(str+k)<=101)  // if input a,b,c
					mx=*(str+k)-87;
				else
					mx=-1;
				if(mx>=0)  // mx is a number for power of 2 
					temp+=(unsigned)pow(2,mx);   // temp is number
				k++;
				
			}

			
			if(temp&board){    // check new input if conflict with old board state
				std::cout<<"Your move invalid!!!"<<std::endl;
				continue;
			}


			for(i=0;i<63;i++)    // check base move data if valid
				if(Move[i]==temp)
				{
					flag=0;   // flag initialize to 0
					break;
				}
			if(flag) std::cout<<"Invalid move!!!"<<std::endl;
		}

		board=temp|board;   // valid move , change the board 
			
		if(GameOver(board))
			gameover=1;
		Display();
	}

}

void Game::Mymove(int &gameover)
{
	int depth=SEARCH_DEPTH;
	int turn = ME;
	int A,B,i;
	int c;
	unsigned myBoard;

	if(!gameover){

		Root->step=0;
		myBoard=board;
		A= -1;
		B= 1;

		// after the AlphaBeta call we can get the best move 
		// in NodeBuff+(Root->bestchild)
		c=AlphaBeta(Root,A,B,depth,turn,myBoard);  
		//c=GameTree(Root,depth,turn,myBoard);
		Temp=NodeBuff+(Root->bestchild);  
#ifdef DEBUG
		if(!memlack)  // if we have enought memory so we can certain
		{
			if(c<0)   
				std::cout<<"I will lose"<<std::endl;
			else if(c>0) 
				std::cout<<" Ha!Ha! I will win"<<std::endl;
			else
				std::cout<<" I don't know who will win"<<std::endl;
		}
		else
				std::cout<<" You have opptunity"<<std::endl;
	
#endif
		std::cout<<"My move is: ";  // show the move
		for(i=0;i<15;i++)
			if((unsigned)power2[i]&Temp->step)
			{
				if(i>=10)
					std::cout<<(char)(i+55);
				else
					std::cout<<i<<" ";
			}

		// update the board state
		Moving(*Temp);
#ifdef DEBUG
		std::cout<<" r= "<<r<<" s ="<<s<<std::endl;  //for testing
#endif
		r=0;
		s=0;
		memlack=0;

		
		if(GameOver(board))
			gameover=1;
		Display();
	}

}

void Game::UnMoving(NODE w,unsigned &myBoard)
{
	myBoard=myBoard&(~w.step);
	w.endnode=0;
}

void Game::Moving(NODE w,unsigned &myBoard)
{

	myBoard=myBoard|w.step;
	if(GameOver(myBoard))
		w.endnode=1;

}

void Game::Moving(NODE w)
{

	board=board|w.step;
	if(GameOver(board))
		w.endnode=1;

}


void Game::FreeNodes()
{
	count=0;
	Root->cnum=0;

}

void Game::Score()
{

}

void Game::GenBaseMove()
{

	int i,j;
	unsigned temp;
	// two token move data
	int temp2[30][2]={
	{0,1},{1,2},{2,3},{3,4},{5,6},{6,7},{7,8},{9,10},{10,11},{12,13},
	{0,5},{5,9},{9,12},{12,14},{1,6},{6,10},{10,13},{2,7},{7,11},{3,8},
	{4,8},{8,11},{11,13},{13,14},{3,7},{7,10},{10,12},{2,6},{6,9},{1,5}};
	// three token move data
	int temp3[18][3]={
		{0,1,2},{1,2,3},{2,3,4},{5,6,7},{6,7,8},{9,10,11},
		{0,5,9},{5,9,12},{9,12,14},{1,6,10},{6,10,13},{2,7,11},
		{4,8,11},{8,11,13},{11,13,14},{3,7,10},{7,10,12},{2,6,9}};

	
	// one token
	for(i=0;i<15;i++)
		Move[i]=(unsigned)power2[i];

	// tow token
	for(i=0;i<30;i++)
	{
		temp=0;
		for(j=0;j<2;j++)
			temp+=(unsigned)power2[temp2[i][j]];
		Move[i+15]=temp;
	}

	// three token
	for(i=0;i<18;i++)
	{
		temp=0;
		for(j=0;j<3;j++)
			temp+=(unsigned)power2[temp3[i][j]];
		Move[i+45]=temp;
	}
}

// for reduce use of the memory 
unsigned Game::Rotate(int angle, unsigned int step)
{
	unsigned temp=0;
	int i;

	switch(angle)
	{
	case 120:
		for(i=0;i<15;i++){
			if((unsigned)power2[i]&step)
				temp+=(unsigned)power2[rotate120[i]];
		}
		break;
	case 240:
		for(i=0;i<15;i++){
			if((unsigned)power2[i]&step)
				temp+=(unsigned)power2[rotate240[i]];
		}
		break;
	case 12090:
		for(i=0;i<15;i++){
			if((unsigned)power2[i]&step)
				temp+=(unsigned)power2[rotate120r[i]];
		}
		break;
	case 24090:
		for(i=0;i<15;i++){
			if((unsigned)power2[i]&step)
				temp+=(unsigned)power2[rotate240r[i]];
		}
		break;
	case 90:
		for(i=0;i<15;i++){
			if((unsigned)power2[i]&step)
				temp+=(unsigned)power2[rotate90[i]];
		}
	}
	return temp;
}

void Game::MyTest()
{
	NODE w;
	w.step=1;
	
	board=(unsigned)pow(2,14)-1;
	board=0;
}

int Game::GameTree(NODE *p, int depth, int turn,unsigned &myBoard)
{
	NODE *q;
	int		c;
	unsigned i;
	unsigned w;
	int temp;

	Moving(*p,myBoard);
	if(p->endnode||depth==0)
	{
		temp=Evaluate(myBoard);
		UnMoving(*p,myBoard);
		if(turn==ME)
			return -temp;
		else
			return temp;

	}
	
	MoveGen(p,turn,myBoard);
	w=p->children;
	for(i=0;i<p->cnum;i++,w++){
		q=NodeBuff+w;
		c=GameTree(q,depth-1,-turn,myBoard);
		switch(turn)
		{
		case ME:
			if(c==1){
				p->bestchild=w;
				return c;
			}
			break;
		case YOU:
			if(c==-1){
				p->bestchild=w;
				return c;
			}
		}
	}
	UnMoving(*p,myBoard);
	return 0;
}

void Game::Reset()
{
	board=0;
	r=0;
	memlack=0;
	count=0;
}

// just for find win pattern using 
// not use in the game program
void Game::FindWin(unsigned Win[],unsigned Loss[],unsigned &iWin,unsigned &iLoss)
{
	unsigned iWinBound=0;
	unsigned  iLossBound=0;
	unsigned j,k,exist;

	iWin=0;
	iLoss=0;

	// left 1 token you will win the first 15 wins
	for(iWin=0;iWin<15;iWin++)
		Win[iWin]=power2[iWin];
	
	while(iWinBound!=iWin)
	{
		//iWinBound=iWin;
	
 		while(iWinBound<iWin)   // looking for lose
		{
			
			MoveGen(Root,ME,Win[iWinBound]);
				
			for(j=1;j<=Root->cnum;j++)  //look all possible move
			{
				

				exist=0;
			
				for(k=0;k<iLoss;k++)  // check if exist
				{
					if(((NodeBuff+j)->step| Win[iWinBound])==Loss[k])
					{
						exist=1;
						break;
					}
				}
				if(!exist)   // if no exist inser to Loss table
				{
					if(iLoss>65534)
					{
						std::cout<<"iloss memory lack"<<std::endl;
						return;
					}
			//		std::cout<<"iLoss="<<iLoss<<std::endl;
					Loss[iLoss++]=((NodeBuff+j)->step) | Win[iWinBound];
				
				}
			}
			FreeNodes();
			iWinBound++;
		}
		

		
		while(iLossBound<iLoss) // looking for win
		{
			MoveGen(Root,ME,Loss[iLossBound]);
			for(j=1;j<=Root->cnum;j++)  //look all possible move
			{
				exist=0;
			
				for(k=0;k<iWin;k++)  // check if exist in Win[]
				{
					if(((NodeBuff+j)->step | Loss[iLossBound])==Win[k])
					{
						exist=1;
						break;
					}
				}
				if(!exist)   // if no exist inser to Loss table
				{
					if(iWin>65534)
					{
						std::cout<<"iWin memory lack"<<std::endl;
						return;
					}
			//		std::cout<<"iWin="<<iWin<<std::endl;
					
					Win[iWin++]=((NodeBuff+j)->step) | Loss[iLossBound];
					if(iWin==6299)
						return;

				}
			
			FreeNodes();
			iLossBound++;
			}
		}
	}
}
