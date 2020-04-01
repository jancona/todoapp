import 'package:flutter/material.dart';
import 'package:todoclient/api.dart';

enum FilterType {
  All,
  Active,
  Completed
}

class ToDoModel extends ChangeNotifier {
  bool _isLoading = true;
  bool get isLoading => _isLoading;
  List<ModelToDo> _toDos = List<ModelToDo>();
  List<ModelToDo> get toDos {
    _toDos.sort((a, b) => a.order.compareTo(b.order));
    return _toDos;
  }
  String _errorMessage; 
  String get errorMessage => _errorMessage;
  FilterType filter = FilterType.All;

  final _apiInstance = TodosApi(ApiClient(basePath: "http://localhost:3000"));

  ToDoModel() { 
    load();
  }

  void load() async {
    _errorMessage = null;
    try {
      _toDos = await _apiInstance.find();
      print("called TodosApi.find");
    } catch (e) {
      _errorMessage ="Exception when calling TodosApi.find: $e";
      print(_errorMessage);
      _toDos = List<ModelToDo>();
    }
    _isLoading = false;
    notifyListeners();
  }
  void add(ModelToDoInput tdi) async {
    _errorMessage = null;
    try {
      var todo = await _apiInstance.addOne(tdi);
      print("called TodosApi.addOne");
      _toDos.add(todo);
    } catch (e) {
      _errorMessage = "Exception when calling TodosApi.addOne: $e";
      print(_errorMessage);
      load();
      return;
    }
    notifyListeners();
  }
  void update(ModelToDo toDo) async {
    _errorMessage = null;
    try {
      var tdi = ModelToDoInput();
      tdi.title = toDo.title;
      tdi.completed = toDo.completed;
      tdi.order = toDo.order;
      var newToDo = await _apiInstance.update(toDo.id, tdi);
      print("called TodosApi.update");
      _toDos[_toDos.indexOf(toDo)] = newToDo;
    } catch (e) {
      _errorMessage = "Exception when calling TodosApi.update: $e";
      print(_errorMessage);
      load();
      return;
    }
    notifyListeners();
  }
  void delete(ModelToDo toDo) async {
    _errorMessage = null;
    try {
      await _apiInstance.delete(toDo.id);
      print("called TodosApi.delete");
      _toDos.remove(toDo);
    } catch (e) {
      _errorMessage = "Exception when calling TodosApi.update: $e";
      print(_errorMessage);
      load();
      return;
    }
    notifyListeners();
  }
}
