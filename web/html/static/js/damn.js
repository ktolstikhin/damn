const allowedGenders = ['male', 'female'];
const minApiCallDelay = 4 * 1000;
const maxApiCallDelay = 8 * 1000;

function getRandomNumber(min, max) {
    return Math.floor(Math.random() * (max - min) + min);
}

function copyToClipboard(text) {
  var dummy = document.createElement('textarea');
  document.body.appendChild(dummy);
  dummy.value = text
  dummy.select();
  document.execCommand('copy');
  document.body.removeChild(dummy);
  $.notify('💾', {
    position: 'top center',
    className: 'success',
  });
}

$(document).ready(function() {

  $('#shareButton').click(function() {
    copyToClipboard(window.location.toString());
  });

  var urlParams = new URLSearchParams(window.location.search)
  var name = decodeURIComponent(urlParams.get('name'));

  if (name) {
    $('#damnName').text(name);
  }

  var gender = urlParams.get('gender');
  if (!allowedGenders.includes(gender)) {
    gender = allowedGenders[0];
  }

  var level = parseInt(urlParams.get('level'));
  if (isNaN(level) || level < 1 || level > 5) {
    level = 1;
  }

  var obscene = urlParams.get('obscene') == '1';

  var callGodDamnApi = function() {
    urlParams = $.param({
      gender: gender,
      level: level,
      obscene: obscene,
    });
    let url = `/api/damn/ru?${urlParams}`;
    $.ajax({
      type: 'GET',
      url: url,
      dataType: 'json',
      contentType: 'application/json',
      success: function(data) {
        damnText = data.tokens.join(' ');
        copyText = name ? `${name} ${damnText}` : damnText;
        var damnItem = `
          <div class="row justify-content-center mb-3 fadeIn">
            <div class="col-md-8 text-center">
              ${damnText}
              <button class="btn" onclick="copyToClipboard('${copyText}')">
                <i class="fa-regular fa-copy copy-icon"></i>
              </button>
            </div>
          </div>
        `;
        $('#damnItems').prepend(damnItem);
        setTimeout(callGodDamnApi, getRandomNumber(minApiCallDelay, maxApiCallDelay));
      },
      error: function(XMLHttpRequest, textStatus, errorThrown) { 
        setTimeout(callGodDamnApi, getRandomNumber(minApiCallDelay, maxApiCallDelay) * 2);
      }
    });
  }

  callGodDamnApi();

});
