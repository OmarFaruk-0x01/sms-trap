from django import forms

class TrapForm(forms.Form):
    phone_number = forms.CharField(
        label="Phone Number",
        widget=forms.NumberInput()
    )
    message = forms.CharField(
        widget=forms.Textarea(),
        label="Message"
    )
    label = forms.ChoiceField(
        choices=[('transactional', 'Transactional'), ('promotional', 'Promotional')],
        label="Select an Option",
        widget=forms.Select()
    )