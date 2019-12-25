$(document).ready( function () {

  const currentUsersTable = $('#currentUsers').DataTable();
  const deletedUsersTable = $('#deletedUsers').DataTable();

  $("#usersForm").submit(function(e){
    e.preventDefault()

    let data = $("#usersForm").serialize()

    $.ajax({
      method: "POST",
      url: "http://localhost:3000/signupaction",
      data: data
    }).done(function( data ) {

      console.log(data);

        // $("#responseMessage").empty()
        // $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');
        //
        // currentCategoriesTable.clear()
        //
        // data.Categories.forEach(function(value){
        //
        //   let rmvButton = '<td><button id="'+value.ID+'" class="btn btn-danger remove">Remove</button></td>'
        //
        //   let array = [value.ID, value.Category, value.Description, rmvButton]
        //
        //   $('#currentCategories').DataTable().row.add(array).draw()
        //
        //     $("#categoriesForm")[0].reset()
        //
        // })



    })

  });//end of submit


  $(document).on("click", ".remove", function(){

    let user_uuid = $(this).attr('id')

    $("#responseMessage").empty();


    bootbox.confirm({
        title: "Destroy User?",
        message: "Do you want to destroy this user?",
        buttons: {
            cancel: {
                label: 'Cancel'
            },
            confirm: {
                label: 'Confirm'
            }
        },
        callback: function (result) {

          if(result){
            $.ajax({
              method: "GET",
              url: "http://localhost:3000/removeUser/"+ user_uuid,
            }).done(function( data ) {

              console.log(data);

                deletedUsersTable.clear()

                if (data.RemoveBy == user_uuid){
                  $("#" + data.RemoveBy).parent().parent().remove()
                  $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

                  data.U.forEach(function(value){

                    let restoreButton = '<td><button id="deleted_'+value.UserUuid+'" class="btn btn-success restore">Restore</button></td>'

                    let array = [ value.Username, value.FirstName, value.LastName, value.Role, restoreButton]

                    $('#deletedUsers').DataTable().row.add(array).draw()

                  })

                }else{
                  $("#responseMessage").append('<h5 class="animated fadeOut">Error! Try Again!</h5>');
                }

                })

          }
        }
    });

    })

    $(document).on("click", ".restore", function(){

      let user_uuid = $(this).attr('id').split("_")

      $("#responseMessage").empty();

      bootbox.confirm({
          title: "Restore User?",
          message: "Do you want to restore this user?",
          buttons: {
              cancel: {
                  label: '<i class="fa fa-times"></i> Cancel'
              },
              confirm: {
                  label: '<i class="fa fa-check"></i> Confirm'
              }
          },
          callback: function (result) {
              if(result){
                $.ajax({
                  method: "GET",
                  url: "http://localhost:3000/restoreUser/"+ user_uuid[1],
                }).done(function( data ) {
                  currentUsersTable.clear()

                  console.log(data);

                  if( data.RemoveBy == user_uuid[1]){
                    $("#deleted_" + data.RemoveBy).parent().parent().remove()
                    $("#responseMessage").append('<h5 class="animated fadeOut">'+data.Message+'</h5>');

                    data.U.forEach(function(value){

                      let rmvButton = '<td><button id="'+value.UserUuid+'" class="btn btn-danger remove">Remove</button></td>'

                      let array = [value.Username, value.FirstName, value.LastName, value.Role, rmvButton]

                      $('#currentUsers').DataTable().row.add(array).draw()

                    })
                  }else{
                    $("#responseMessage").append('<h5 class="animated fadeOut">Error! Try Again!</h5>');
                  }



                    })
              }
          }
      });


      })


});//end of doc ready
