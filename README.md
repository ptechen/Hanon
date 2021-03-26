# Hanon

### Default execution order
    vec![String::from("selects"),
        String::from("each"),
        String::from("select_params"),
        String::from("nodes"),
        String::from("has"),
        String::from("contains")];
        
    selects > each > (one or all or fields) > ... text_attr_html > (text or attr or html);
    selects > select_params > selects > ... text_attr_html > (text or attr or html);
    selects > nodes > has > contains > text_attr_html > (text or attr or html);


### Support:
| Capricorn | support | example |val type|
| :----: | :----: | :----- |:----:|
| selects element | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name | String |
| selects class | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - .class_name | String | 
| selects class element | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - .class_name <br> &nbsp; &nbsp; &nbsp; - element_name | String | 
| nodes first | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; nodes: <br> &nbsp; &nbsp; &nbsp; first: true | String | 
| nodes last | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; nodes: <br> &nbsp; &nbsp; &nbsp; last: true | String | 
| nodes eq | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; nodes: <br> &nbsp; &nbsp; &nbsp; eq: 0 | String | 
| nodes parent | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; nodes: <br> &nbsp; &nbsp; &nbsp; parent: true | String | 
| nodes children | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; nodes: <br> &nbsp; &nbsp; &nbsp; children: true | String | 
| nodes prev_sibling | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; nodes: <br> &nbsp; &nbsp; &nbsp; prev_sibling: true | String | 
| nodes next_sibling | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; nodes: <br> &nbsp; &nbsp; &nbsp; next_sibling: true | String | 
| each one | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; each: <br> &nbsp; &nbsp; &nbsp; one: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; selects:<br>&nbsp; &nbsp; &nbsp;  &nbsp; &nbsp; &nbsp; &nbsp; - .class_name<br>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;  ... | String | 
| each all | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; each: <br> &nbsp; &nbsp; &nbsp; all: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; selects:<br>&nbsp; &nbsp; &nbsp;  &nbsp; &nbsp; &nbsp; &nbsp; - .class_name<br>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;  ... | Array | 
| each fields | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; each: <br> &nbsp; &nbsp; &nbsp; fields: <br> &nbsp; &nbsp; &nbsp; &nbsp; field_name: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; selects:<br>&nbsp; &nbsp; &nbsp;  &nbsp; &nbsp; &nbsp; &nbsp; - .class_name<br>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;  ... <br> &nbsp; &nbsp; &nbsp; &nbsp; field_name1: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; selects:<br>&nbsp; &nbsp; &nbsp;  &nbsp; &nbsp; &nbsp; &nbsp; - .class_name<br>&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;  ...  | Map | 
| select_params | ✔ | field_name: <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; select_params: <br> &nbsp; &nbsp; &nbsp; selects:<br>&nbsp; &nbsp; &nbsp;  &nbsp; &nbsp; - .class_name<br>&nbsp; &nbsp; &nbsp;  ... | ... | 
| text | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; text_attr_html: <br> &nbsp; &nbsp; &nbsp; text: true | String |
| attr | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; text_attr_html: <br> &nbsp; &nbsp; &nbsp; attr: true | String |
| html | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; text_attr_html: <br> &nbsp; &nbsp; &nbsp; html: true | String |
| contains text | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; text: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - test | String |
| contains html | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; html: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - test | String |
| contains attr | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; attr: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - test | String |
| contains class | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; class: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - test | String |
| not_contains text | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; not_contains: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; text: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - test | String |
| not contains html | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; contains: <br> &nbsp; &nbsp; &nbsp; not_contains: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; html: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - test | String |
| exec order | ✔ | field_name:<br> &nbsp; exec_order: <br> &nbsp; &nbsp; &nbsp; - selects <br> &nbsp; &nbsp; &nbsp; - has <br> &nbsp; &nbsp; &nbsp; - nodes <br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; has: <br> &nbsp; &nbsp; &nbsp; class: class_name <br> &nbsp; nodes: <br> &nbsp; &nbsp; &nbsp; first: true | String |
| data format splits | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; data_format: <br> &nbsp; &nbsp; &nbsp; splits: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - { key: str } | Array |
| data format splits | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; data_format: <br> &nbsp; &nbsp; &nbsp; splits: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - { key: str, index: 0 } | String |
| data format replaces | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; data_format: <br> &nbsp; &nbsp; &nbsp; replaces: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - str | String |
| data format deletes | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; data_format: <br> &nbsp; &nbsp; &nbsp; deletes: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - str | String |
| data format find | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; data_format: <br> &nbsp; &nbsp; &nbsp; find: <br> &nbsp; &nbsp; &nbsp; &nbsp; - regex | String |
| data format find_iter | ✔ | field_name:<br> &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; - element_name <br> &nbsp; data_format: <br> &nbsp; &nbsp; &nbsp; find_iter: <br> &nbsp; &nbsp; &nbsp; &nbsp; - regex | Array |
| Multi-version regular matching err | ✔ |regexes_match_parse_html: <br>  &nbsp; &nbsp; - regex: regex <br> &nbsp; &nbsp; &nbsp; version: 1 <br> &nbsp; &nbsp; &nbsp; err: err_msg | Err |
| Multi-version regular matching fields | ✔ |regexes_match_parse_html: <br>  &nbsp; &nbsp; - regex: regex <br> &nbsp; &nbsp; &nbsp; version: 1 <br> &nbsp; &nbsp; &nbsp; fields: <br> &nbsp; &nbsp; &nbsp; &nbsp; field_name: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; ... <br> &nbsp; &nbsp; &nbsp; &nbsp; field_name: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; selects: <br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; ... | Map |
