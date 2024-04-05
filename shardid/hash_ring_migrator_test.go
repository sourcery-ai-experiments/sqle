package shardid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHRMigrator(t *testing.T) {

	//	2 dbs   ->   3 dbs  -> data
	//  -> 2439456            1149
	// 46916880 E0 0  !
	// 63694499 E1 1
	//	<-	80472118 E2 2
	//	<-  84017712 S2 2
	//  ->  111074370         638
	// 117572950 S0 0 !
	// 134350569 S1 1
	//	<-  214987260 G2 2
	// 248542498 G0 0 !
	// 265320117 G1 1
	// 316638712 M0 0
	// 333416331 M1 1
	//	<-  350193950 M2 2
	//	<-  351179688 K2 2
	// 384734926 K0 0 !
	// 401512545 K1 1
	//	<-  484709092 O2 2
	// 518264330 O0 0 !
	// 535041949 O1 1
	//	<-  2228889920 C2 2
	// 2262445158 C0 0 !
	// 2279222777 C1 1
	// 2330541372 I0 0
	// 2347318991 I1 1
	//	<-  2364096610 I2 2
	// 2597703348 A0 0 !
	// 2600263204 Q0 0
	// 2614480967 A1 1
	// 2617040823 Q1 1
	//	<-  2631258586 A2 2
	//	<-  2633818442 Q2 2
	//  -> 4113327457          150

	r1 := NewHR(2, WithReplicas(defaultReplicas...))
	r2 := NewHR(3)

	m := NewHRMigrator(r1, r2)

	vn := map[uint32]bool{
		46916880:   true, // E0
		117572950:  true, // S0
		248542498:  true, // G0
		384734926:  true, // K0
		518264330:  true, // O0
		2262445158: true, // C0
		2597703348: true, // A0
	}
	require.Equal(t, vn, m.affectedVNodes)
	require.Equal(t, []int{0}, m.affectedDbs)

	require.True(t, m.Has("1149")) // < E0 ! first node
	require.False(t, m.Has("E0"))  // == E0 !
	require.True(t, m.Has("E1"))   // == E1
	require.True(t, m.Has("638"))  // E1 < 638 < S0
	require.False(t, m.Has("S0"))  // == S0 !
	require.False(t, m.Has("C0"))  // == C0 !
	require.False(t, m.Has("C1"))  // == C1
	require.True(t, m.Has("150"))  // > Q1 last node
}