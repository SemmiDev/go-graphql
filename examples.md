**create a user andget a token and apply in http header for others request
**
```
mutation {
	createUser(input: {username: "user1", password: "123"})
 }
```


**login**
```
mutation {
	login(input: {username: "user1", password: "123"})
 }
```

**create a link**

```
mutation {
	createLink(input: {title: "real link!", address: "www.graphql.org"}){
		 user{
      			name
    		}
  	}
 }
```

**generate new token**
```
mutation {
  refreshToken(input: {token: "tokennnya"})
}
```


**get links**
```
query {
  links {
    title
    address
    id
    user {
      id
      name
    }
  }
}
```


**create an author**
```
mutation {
  createAuthor(firstName: "sammi", lastName: "dev"){
    id
    firstName
    lastName
  }
}
```

**create a book**
```
mutation {
  createBook(title: "belajar go", Author: "1"){
    id
    title
    Author {
      id
      firstName
      lastName
    }
  }
}
```

**book by id**
```
query {
  bookByID(id: "1") {
    title
    Author {
      firstName
      lastName
    }
  }
}
```

**all books**
```
query {
  allBooks {
    id
    title
    Author {
      firstName
      lastName
    }
  }
}
```


**author by id**
```
query {
  authorByID(id: "1") {
    firstName
    lastName
  }
}
```


**all authors**
```
query {
  allAuthors {
    firstName
    lastName
  }
}
```