# active-standby
zookeeper를 사용하여 active-standby 이중화 구성 테스트

## zookeeper usecase

### Barriers

- Kafka에서 Broker중 Controller를 선출할 때 사용하는 방식
- 가장 먼저 /Controller를 선점하는 Broker가 Controller의 역할 수행
- /Master 노드를 선점하는 프로세스가 Master 역할을 하도록 구성해보자.

#### 문제점
![스크린샷 2020-12-28 오후 6 18 34](https://user-images.githubusercontent.com/22065697/103206879-463bf080-4940-11eb-837d-dd0418067dbc.png)

- Standby 3개 프로세스가 모두 Master가 된다. Lock을 도입해봐야 할듯

### Distribute Lock 도입

![ezgif com-gif-maker (1)](https://user-images.githubusercontent.com/22065697/103206901-5522a300-4940-11eb-8d30-67f74b28dfba.gif)

- Active -> Standby failover가 정상적으로 수행된다.

### Leader Election

- Active-Standby 총 2대의 서버로 이중화 되기 대문에 Leader Election 방법을 사용하기엔 적합하지 않다.
