const guessTextBox = document.querySelector('input[name="guess"]');
const answerDiv = document.querySelector("#answer");

function Setup(word){
    document.querySelector("#guessButton").addEventListener('click', function guess(){
        if(guessTextBox.value !== word){
            guessTextBox.setCustomValidity("Your guess is not correct");
            guessTextBox.reportValidity();
        }
        else{
            answerDiv.classList.remove('hidden')
        }
    });
    
    document.querySelector("#toggleButton").addEventListener('click', function guess(){
        answerDiv.classList.toggle('hidden')
    });
}

