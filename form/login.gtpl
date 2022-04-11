<html>
  <head>
    <title></title>
  </head>
  <body>
    <form action="/login" method="post">
      ユーザ名:<input type="text" name="username">
      メールアドレス:<input type="text" name="email">
      年齢:<input type="number" name="age">
      パスワード:<input type="text" name="password">
      <input type="hidden" name="token" value="{{.}}">
      <input type="submit" value="ログイン">
    </form>
  </body>
</html>