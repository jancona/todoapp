part of todoclient.api;

class ModelToDo {
  
  bool completed = null;
  
  String id = null;
  
  int order = null;
  
  String title = null;
  
  String url = null;
  ModelToDo();

  @override
  String toString() {
    return 'ModelToDo[completed=$completed, id=$id, order=$order, title=$title, url=$url, ]';
  }

  ModelToDo.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    completed = json['completed'];
    id = json['id'];
    order = json['order'];
    title = json['title'];
    url = json['url'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (completed != null)
      json['completed'] = completed;
    if (id != null)
      json['id'] = id;
    if (order != null)
      json['order'] = order;
    if (title != null)
      json['title'] = title;
    if (url != null)
      json['url'] = url;
    return json;
  }

  static List<ModelToDo> listFromJson(List<dynamic> json) {
    return json == null ? List<ModelToDo>() : json.map((value) => ModelToDo.fromJson(value)).toList();
  }

  static Map<String, ModelToDo> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ModelToDo>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ModelToDo.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ModelToDo-objects as value to a dart map
  static Map<String, List<ModelToDo>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ModelToDo>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ModelToDo.listFromJson(value);
       });
     }
     return map;
  }
}

