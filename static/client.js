// Form submission with fetch API
document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('horoscope-form');
    if (!form) return;

    form.addEventListener('submit', function(e) {
        console.log('Form submit triggered');
        const formData = new FormData(form);
        console.log('Form data:', Object.fromEntries(formData));
        e.preventDefault();

        const formData = new FormData(form);
        const data = {};
        formData.forEach(function(value, key) {
            data[key] = value;
        });

        // Number conversions
        const day = parseInt(data.day);
        const month = parseInt(data.month);
        const year = parseInt(data.year);
        const hour = parseInt(data.hour);
        const gender = parseInt(data.gender);

        if (!day || !month || !year || !hour || !gender) {
            alert('Vui lòng điền đầy đủ thông tin!');
            return;
        }

        data.day = day;
        data.month = month;
        data.year = year;
        data.hour = hour;
        data.gender = gender;

        // Call API endpoint - POST to /api/stars
        fetch('/api/stars', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        })
        .then(response => {
            if (!response.ok) throw new Error('API error: ' + response.status);
            return response.text();
        })
        .then(html => {
            console.log('API response received:', html.substring(0, 200) + '...');
            const chartOutput = document.getElementById('chart-output');
            if (chartOutput) {
                chartOutput.innerHTML = html;
                chartOutput.classList.remove('hidden');
                console.log('Chart output updated');
            }
        })
        .catch(error => {
            console.error('Error:', error);
            hideLoading(true);
            alert('Đã xảy ra lỗi khi gọi API. Vui lòng thử lại.');
        });
    });
});