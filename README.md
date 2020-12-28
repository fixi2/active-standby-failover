# active-standby
zookeeper를 사용하여 active-standby 이중화 구성 테스트

## zookeeper usecase

### Barriers

- Kafka에서 Broker중 Controller를 선출할 때 사용하는 방식
- 가장 먼저 /Controller를 선점하는 Broker가 Controller의 역할 수행
- /Master 노드를 선점하는 프로세스가 Master 역할을 하도록 구성해보자.

### 문제점
<img src="https://media.oss.navercorp.com/user/16779/files/4389cd00-4939-11eb-9cdb-069fe13c8404" width="50%" >

- Standby 3개 프로세스가 모두 Master가 된다. Lock을 도입해봐야 할듯

### Leader Election

- Active-Standby 총 2대의 서버로 이중화 되기 대문에 Leader Election 방법을 사용하기엔 적합하지 않다.