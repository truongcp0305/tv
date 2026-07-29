package templates

// ResetGenderButtonsJS contains the JavaScript code for resetting gender button states.
const ResetGenderButtonsJS = `
function resetGenderButtons() {
    // Reset all gender buttons to default state
    document.querySelectorAll('.gender-btn').forEach(btn => {
        btn.className = 'gender-btn flex-1 rounded-lg border border-slate-300 bg-white py-3 text-center text-sm font-medium text-slate-700 hover:border-indigo-500 hover:bg-indigo-50 hover:text-indigo-600 transition-all focus:outline-none focus:ring-2 focus:ring-indigo-200';
    });
    
    // Reset gender value to male (1) by default
    document.getElementById('gender').value = '1';
    document.getElementById('btn-male').classList.add('bg-indigo-50', 'border-indigo-500', 'text-indigo-600');
}

document.addEventListener('DOMContentLoaded', function() {
    // Initialize first gender button as selected
    const maleBtn = document.getElementById('btn-male');
    if (maleBtn) {
        maleBtn.classList.add('bg-indigo-50', 'border-indigo-500', 'text-indigo-600');
    }
    
    // Handle gender button clicks
    document.querySelectorAll('.gender-btn').forEach(btn => {
        btn.addEventListener('click', function() {
            document.querySelectorAll('.gender-btn').forEach(b => {
                b.className = 'gender-btn flex-1 rounded-lg border border-slate-300 bg-white py-3 text-center text-sm font-medium text-slate-700 hover:border-indigo-500 hover:bg-indigo-50 hover:text-indigo-600 transition-all focus:outline-none focus:ring-2 focus:ring-indigo-200';
            });
            
            this.className = 'gender-btn flex-1 rounded-lg border border-indigo-500 bg-indigo-50 py-3 text-center text-sm font-medium text-indigo-600 transition-all focus:outline-none focus:ring-2 focus:ring-indigo-200';
            
            // Update hidden input value
            const genderValue = this.textContent.trim() === 'Nam' ? '1' : '-1';
            document.getElementById('gender').value = genderValue;
        });
    });
    
    // Calendar toggle logic
    const calendarToggle = document.getElementById('calendar-toggle');
    const calendarMode = document.getElementById('calendarMode');
    if (calendarToggle && calendarMode) {
        calendarToggle.addEventListener('click', function() {
            const isLunar = calendarMode.value === 'lunar';
            calendarMode.value = isLunar ? 'solar' : 'lunar';
            this.textContent = isLunar ? 'Dương lịch' : 'Âm lịch';
        });
    }
});
`
