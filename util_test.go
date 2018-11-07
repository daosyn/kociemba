package kociemba

import "testing"

func TestConvert(t *testing.T) {
	convert("test")
}

func TestCreateTwistMoveTable(t *testing.T) {
	createTwistMoveTable()
}

/*

TwistMove: array[0..2187-1,Ux1..Bx3] of Word;

procedure CreateTwistMoveTable;
var c: CubieCube; i,k: Integer; j: TurnAxis;
begin
  c:= CubieCube.Create;//create a cube c on the cubie level
  for i:=0 to 2187-1 do
  begin
    c.InvCornOriCoord(i);//generate a permutation with corner orientation i
    for j:= U to B do
    begin
      for k:= 0 to 3 do //k=3 restores the original state
      begin
        c.Move(j);//apply all 18 face turns on c
        if k<>3 then TwistMove[i,Move(3*Ord(j)+k)]:=c.CornOriCoord;//save result in the array
      end;
    end;
  end;
  c.Free;
end;
*/
