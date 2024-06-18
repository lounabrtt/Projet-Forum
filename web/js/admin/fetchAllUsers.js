function fetchAllUsers() {
  fetch("/api/users")
    .then((response) => response.json())
    .then((data) => {
      const tbody = document.querySelector("tbody");
      tbody.innerHTML = "";

      data.forEach((user) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${user.uuid}</td>
            <td>${user.lastName}</td>
            <td>${user.firstName}</td>
            <td>${user.birthdate}</td>
            <td>${user.email}</td>
            <td>
              <form action="/update-role" method="post">
                <select name="role">
                  <option value="visitor">Visitor</option>
                  <option value="user">User</option>
                  <option value="moderator">Moderator</option>
                  <option value="admin">Admin</option>
                </select>
                <input type="hidden" name="uuid" value="${user.uuid}" />
                <button type="submit">Update</button>
              </form>
            </td>
            <td><button onclick="deleteUser('${user.uuid}')">Delete</button></td>
          `;
        tbody.appendChild(row);
      });
    })
    .catch((error) => console.error("Error fetching users:", error));
}

fetchAllUsers();
