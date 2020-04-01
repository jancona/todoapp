// Copyright 2018 The Flutter team. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:todoclient/api.dart';
import 'model.dart';

void main() => runApp(ChangeNotifierProvider(
      create: (context) => ToDoModel(),
      child: ToDoApp(),
    ));

class ToDoApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter • TodoMVC',
      home: ToDos(),
    );
  }
}

class ToDos extends StatelessWidget {
  final _biggerFont = const TextStyle(fontSize: 18.0);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text('Todos'),
        ),
        body: Column(
          children: <Widget>[ToDoForm(), _buildTodos()],
        ));
  }

  Widget _buildTodos() {
    return Consumer<ToDoModel>(
      builder: (context, model, child) {
        if (model.isLoading) {
          return CircularProgressIndicator();
        }
        if (model.errorMessage != null) {
          WidgetsBinding.instance.addPostFrameCallback((_) {
            Scaffold.of(context)
                .showSnackBar(SnackBar(content: Text(model.errorMessage)));
          });
        }
        return Expanded(
            child: Scrollbar(
                child: ListView.builder(
                    shrinkWrap: true,
                    padding: const EdgeInsets.all(16.0),
                    itemCount: 2 * model.toDos.length,
                    itemBuilder: (context, i) {
                      if (i.isOdd) return Divider();
                      final index = i ~/ 2;
                      var t = model.toDos[index];
                      return _buildRow(model, t);
                    })));
      },
    );
  }

  Widget _buildRow(ToDoModel model, ModelToDo todo) {
    // final bool alreadySaved = _saved.contains(todo);
    return ListTile(
      leading: Icon(
        todo.completed ? Icons.check_circle : Icons.check_circle_outline,
        color: todo.completed ? Colors.red : null,
      ),
      title: Text(
        todo.title,
        style: _biggerFont,
      ),
      onTap: () {
        todo.completed = !todo.completed;
        model.update(todo);
      },
    );
  }
}

class ToDoForm extends StatefulWidget {
  @override
  ToDoFormState createState() {
    return ToDoFormState();
  }
}

// Define a corresponding State class.
// This class holds data related to the form.
class ToDoFormState extends State<ToDoForm> {
  // Create a global key that uniquely identifies the Form widget
  // and allows validation of the form.
  //
  // Note: This is a `GlobalKey<FormState>`,
  // not a GlobalKey<ToDoFormState>.
  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    final _tc = TextEditingController();
    final _model = Provider.of<ToDoModel>(context, listen: false);
    // Build a Form widget using the _formKey created above.
    return Form(
        key: _formKey,
        child: Row(children: <Widget>[
          Expanded(
              child: Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: TextFormField(
                    controller: _tc,
                    // The validator receives the text that the user has entered.
                    validator: (value) {
                      if (value.isEmpty) {
                        return 'Please enter some text';
                      }
                      return null;
                    },
                  ))),
          Padding(
              padding: const EdgeInsets.symmetric(vertical: 16.0),
              child: RaisedButton(
                onPressed: () {
                  // Validate returns true if the form is valid, otherwise false.
                  if (_formKey.currentState.validate()) {
                    var tdi = ModelToDoInput();
                    tdi.title = _tc.text;
                    _model.add(tdi);
                    _tc.clear();
                  }
                },
                child: Text('Submit'),
              ))
        ]));
  }
}
