# 4

## sns.go
#### How to execute
In the same directory with **links.txt** and **nicknames.txt**, input `go run sns.go`.

#### It includes 
- A function to count step from one node to another (**bfs**)

_- execution result -_

![Screenshot 2019-06-19 at 15 08 03](https://user-images.githubusercontent.com/34668695/59740801-1e2bf580-92a4-11e9-9b4b-29db68e23d16.png)

- A function to display if the other nodes are connected from one node (**showAllSteps**)

_- execution result -_

![Screenshot 2019-06-20 at 16 16 45](https://user-images.githubusercontent.com/34668695/59828873-e93ca300-9376-11e9-829a-0523f70327e2.png)
(omitting the rest)

- and small tests.
  
## wiki.go
#### How to execute
In the same directory with **wiki_links.txt** (originally links.txt) and **wiki_pages.txt** (originally pages.txt), input

`go run wiki.go`.

 - It search the shortest route and count steps from one node to another.
 
_- execution result -_

![Screenshot 2019-06-20 at 19 29 18](https://user-images.githubusercontent.com/34668695/59843814-498e0d80-9394-11e9-8326-cd8c987dec42.png)

![Screenshot 2019-06-20 at 19 29 05](https://user-images.githubusercontent.com/34668695/59843815-4a26a400-9394-11e9-85b9-ef4e437fd2b9.png)

![Screenshot 2019-06-20 at 19 28 49](https://user-images.githubusercontent.com/34668695/59843816-4a26a400-9394-11e9-927d-2fa76d4eb1a3.png)

![Screenshot 2019-06-20 at 19 35 16](https://user-images.githubusercontent.com/34668695/59844190-30d22780-9395-11e9-9422-295e16f92366.png)

## wiki_isolated.go
- It shows pages that doesn't have connection with the other ones.

_- execution result（抜粋） -_

![Screenshot 2019-06-20 at 19 16 56](https://user-images.githubusercontent.com/34668695/59844363-afc76000-9395-11e9-89e0-4e9a9d26f641.png)

- It shows 1438 isolated pages, but actually that pages have links in their real page.

![Screenshot 2019-06-20 at 20 02 30](https://user-images.githubusercontent.com/34668695/59844636-5e6ba080-9396-11e9-8f5c-4b69f1001e89.png)

