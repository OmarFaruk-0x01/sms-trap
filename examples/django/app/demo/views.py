from django.shortcuts import render, redirect
from .forms import TrapForm
from sms import send_sms, get_sms_provider


def home(request):
    form = TrapForm()
    print(1)
    return render(request, 'home.html', {'form': form})


def submit_form(request):
    if request.method == 'POST':
        form = TrapForm(request.POST)
        if form.is_valid():
            phone_number = form.cleaned_data['phone_number']
            message = form.cleaned_data['message']
            label = form.cleaned_data['label']

            send_sms([phone_number], message, {
                'label': label
            })

    return redirect('/')
