document.addEventListener('DOMContentLoaded', () => {
    fetch('http://localhost:8080/api')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            if (data.message) {
                document.getElementById('api-response').textContent = data.message;
            } else {
                document.getElementById('api-response').textContent = 'No message property found in response.';
            }
        })
        .catch(error => {
            console.error('error-->', error);
            document.getElementById('api-response').textContent = 'Failed to fetch data from the server.';
        });
});