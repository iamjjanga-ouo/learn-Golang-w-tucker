## for문
- `continue`, `break` 다 지원

### 이중 for문의 탈출조건
1. 플래그 변수 사용
   ```Go
   found := false
   for i := 0; i < 5; i++ {
     for j := 0; j <4; j++ {
       if i*j == 4 {
         found = true
         break
       }
     }
     if found == true {
       break
     }
   }
   ```
2. 레이블 사용
  ```Go
  OuterFor:
    for i := 0; i < 5; i++ {
      for j := 0; j <4; j++ {
        if i*j == 4 {
          break OuterFor // 레이블에 가장 먼저 포함된 for문까지 종료
        }
      }
    }
  ```
- 편리한 방법이나, 혼동을 불러올 수 있는 코드
- 꼭 필요한 경우에만 사용하면 좋을듯
- 클린 코드를 지향하려면 중첩된 내부 로직은 -> 함수로 묶어 중첩을 줄이고, 플래그 및 레이블 사용을 최소화