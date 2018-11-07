package kociemba

func convert(cubestring string) [54]Facelet {
	var cube = [54]Facelet{}
	for i := 0; i < 54; i++ {
		cube[i] = Facelet(i)
	}
	return cube
}

func createTwistMoveTable() {

	// TODO create cubie cube
	for i := 0; i < maxTwists; i++ {
		// TODO generate a permutation with corner orientation i
		// for f := range Face {

		// }
	}
}

/*
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
