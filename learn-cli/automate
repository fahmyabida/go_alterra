#how to use --> sh automate {param1} {param2} {param3} 
#!/bin/sh
date="$(date)"
folder_name="$1 $date"
mkdir "$folder_name"
mkdir "$folder_name/about_me"
mkdir "$folder_name/about_me/personal"
mkdir "$folder_name/about_me/professional"
mkdir "$folder_name/my_friends"
mkdir "$folder_name/my_system_info"

#facebook.txt dan linkedin.txt
facebook_path="$folder_name/about_me/personal/facebook.txt"
linkedin_path="$folder_name/about_me/professional/linkedin.txt"
touch "$facebook_path" "$linkedin_path"
echo "https://www.facebook.com/$2" >> "$facebook_path"
echo "https://www.linkedin.com/in/$3" >> "$linkedin_path"

#list_of_my_friends.txt
list_of_my_friends_path="$folder_name/my_friends/list_of_my_friends.txt"
curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt --output "$list_of_my_friends_path"

#about_this_laptop.txt
about_this_laptop_path="$folder_name/my_system_info/about_this_laptop.txt"
touch "$about_this_laptop_path"
echo "My username: $(uname -n)" >> "$about_this_laptop_path"
echo "With host: $(uname -a)" >> "$about_this_laptop_path"

#internet_connection.txt
internet_connection_path="$folder_name/my_system_info/internet_connection.txt"
ping -c 3 google.com >> "$internet_connection_path"