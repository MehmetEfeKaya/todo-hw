document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('todo-form');
    const todoInput = document.getElementById('new-todo');
    const todoList = document.getElementById('todos');

    form.addEventListener('submit', function(e) {
        e.preventDefault();
        const task = todoInput.value.trim();
        if (task) {
            fetch('/todos', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ task }),
            }).then(response => {
                if (response.ok) {
                    loadTodos();
                    todoInput.value = '';
                }
            });
        }
    });

    todoList.addEventListener('click', function(e) {
        const id = e.target.parentElement.dataset.id;

        if (e.target.classList.contains('delete')) {
            fetch(`/todos/${id}`, {
                method: 'DELETE',
            }).then(response => {
                if (response.ok) {
                    loadTodos();
                }
            });
        }

        if (e.target.classList.contains('complete')) {
            fetch(`/todos/${id}/complete`, {
                method: 'PUT',
            }).then(response => {
                if (response.ok) {
                    loadTodos();
                }
            });
        }
    });

    function loadTodos() {
        fetch('/todos')
            .then(response => response.json())
            .then(todos => {
                todoList.innerHTML = '';
                todos.forEach(todo => {
                    const li = document.createElement('li');
                    li.dataset.id = todo.id;
                    li.innerHTML = `
                        <span class="${todo.completed ? 'completed' : ''}">${todo.task}</span>
                        <button class="complete">${todo.completed ? 'Uncomplete' : 'Complete'}</button>
                        <button class="delete">Delete</button>
                    `;
                    todoList.appendChild(li);
                });
            });
    }

    loadTodos();
});
