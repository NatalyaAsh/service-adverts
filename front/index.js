const todosNode = document.querySelector('.js-todos');
//const inputNode = document.querySelector('.js-input');
const btnNode = document.querySelector('.js-btn');
let todos = [];

//function addTodo(text) {
function addUser(name, email, phon) {  
  const todo = {
    //text: name+email+phon,
    //text,
    name: name,
    email: email,
    phon: phon,
    done: false,
    id: `${Math.random()}`
  }

  const user = {
    name: name,
    email: email,
    phon: phon
  }

  todos.push(todo);

  //const urlAdvert = 'http://localhost:8080/advert';
  const urlUser = 'http://localhost:8080/user';
  

  fetch(urlUser, {
    method: 'POST',
    //method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(user)
  })
  .then(response => {
    if (!response.ok) {
      throw new Error('Network response was not ok ' + response.statusText);
    }
    return response.json();
  })
  .then(todo => {
    console.log('Успех:', todo);
  })
  .catch(error => {
    console.error('Ошибка:', error);
  });
}

function deleteTodo(id) {
  todos.forEach(todo => {
    if (todo.id === id) {
      todo.done = true;
    }
  })
}

function render() {
  console.log(todos);
  let html = ``;

  todos.forEach(todo => {
    if (todo.done) {
      return;
    }

    html += `
      <div>
        ${todo.name}
        ${todo.email}
        ${todo.phon}
        <button data-id='${todo.id}'>Сделано</button>
      </div><br>
    `;
  })

  todosNode.innerHTML = html;
}

//addTodo(`Купить хлеба`);
//addTodo(`Купить молока`);

btnNode.addEventListener('click', () => {
  //const text = inputNode.value;
  const name = document.getElementById("todo").value;
  const email = document.getElementById("email").value;
  const phon = document.getElementById("phon").value;
  //addTodo(text+' '+email+' '+phon);
  addUser(name,email,phon);

  render();
});

todosNode.addEventListener('click', (event) => {
  if (event.target.tagName !== 'BUTTON') {
    return;
  }

  const id = event.target.dataset.id;

  deleteTodo(id);
  render();
})

render();