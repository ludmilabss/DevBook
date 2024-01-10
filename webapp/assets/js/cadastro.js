$('#formulario-cadastro').on('submit', criarUsuario);


function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }
    }).done(function() { // 200 201 204
        alert("Usuário cadastrado com sucesso!");
    }).fail(function(erro) { // 400 404 401 403 500
        console.log(erro)
        alert("Erro ao cadastrar o usuário!");
    });
}