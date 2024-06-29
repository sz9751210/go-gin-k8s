db = connect("mongodb://localhost:27017/devops");

var initialData = [
    {
        link: "http://example.com",
        name: "Tom",
        username: "No. 189, Grove St, Los Angeles",
        password: "aaaaaaaaaaa",
    },
    {
        link: "http://example.com",
        name: "Jerry",
        username: "No. 123, Elm St, New York",
        password: "123456",
    }
];

db.linklist.insertMany(initialData);

print("Initial data inserted successfully");
