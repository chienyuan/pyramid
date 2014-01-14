// Game.h: interface for the Game class.
//
//////////////////////////////////////////////////////////////////////

#if !defined(AFX_GAME_H__22E0F382_9E01_11D3_A09B_0000C0B175DD__INCLUDED_)
#define AFX_GAME_H__22E0F382_9E01_11D3_A09B_0000C0B175DD__INCLUDED_

#if _MSC_VER > 1000
#pragma once
#endif // _MSC_VER > 1000

#define ME 1
#define YOU -1

//#define DEBUG 
#define MAX_SIZE_MEM 20000
#define SEARCH_DEPTH 5

struct NODE
{
	unsigned	step;
	unsigned	bestchild;
	unsigned	children;
	unsigned	cnum;
	int			endnode;
	int			value;
};




class Game  
{
protected:
	NODE		*Root,*Temp,*NodeBuff;
	unsigned    Win[6299];
	int			id,number;
	unsigned	count;
	unsigned	board;
	unsigned	Move[64];
	int			rotate120[15],rotate240[15],rotate90[15];
	int			rotate120r[15],rotate240r[15];
	unsigned	power2[15];
public:
	void FindWin(unsigned Win[],unsigned Loss[],unsigned&,unsigned&);
	void Reset();
	int s;
	int memlack;
	unsigned long r;
	int GameTree(NODE *p,int depth,int turn,unsigned &myBoard);
	void MyTest();
	unsigned Rotate(int angle,unsigned step);
	int RotateTest(NODE *p,unsigned move,unsigned myBoard);


	void GenBaseMove();
	void Score();
	void FreeNodes();
	void Moving(NODE w,unsigned &myBoard);
	void Moving(NODE w);

	void UnMoving(NODE w,unsigned &myBoard);
	void Mymove(int &gameover);
	void Yourmove(int &gameover);
	int IfGameOver(NODE w,unsigned myBoard);
	int GameOver(unsigned myBoard);
	void Display();
	void MoveGen(NODE *p,int turn,unsigned myBoard);
	void UnMoveGen(NODE *p, int turn,unsigned myBoard);

	int Evaluate(unsigned myBoard);
	int AlphaBeta(NODE *p,int a,int b,int depth,int turn,unsigned myBoard);
	
	Game();
	virtual ~Game();

};

#endif // !defined(AFX_GAME_H__22E0F382_9E01_11D3_A09B_0000C0B175DD__INCLUDED_)
