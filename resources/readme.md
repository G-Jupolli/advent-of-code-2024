## Resources

This dir is for the input files for the questions. \
All files **must** follow this format:

```
day_{day}_data.txt

e.g. day_1_data.txt
```

Not including the data when trying to get data for a day will trigger an exception. \
You can opptionally use `small` instead of `data` to use a smaller dataset. \
To do this, in a day's mod you will need to use `GetFile(day, false)`.
