* json 格式就是兩種
* 一種就是 key value
  * "class": "class a"
* 一種就是 slice
  * "skill": ["javascript", "golang"]
* 要被 unmarshal 的 struct 裡面的 key 或是 field 一定要大寫，才會被 json 這個包所讀到（超級重要）
