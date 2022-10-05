
from bs4 import BeautifulSoup
import requests
from IPython.core.display import display,HTML
import time

print('Filter out skills which you do not know')
unfamiliar_skills = list(map(str,input('-->').split(',')))
print(f'Filtering out {unfamiliar_skills}')
# --------------------------------------------------------------
print('Type in the skills you know')
familiar_skills = list(map(str,input('-->').split(',')))
print(f'Filtering in {familiar_skills}')

# def find_jobs():
html_text = requests.get('https://www.timesjobs.com/candidate/job-search.html?searchType=personalizedSearch&from=submit&txtKeywords=Analysis&txtLocation=').text
soup = BeautifulSoup(html_text,'lxml')

jobs = soup.find_all('li',class_="clearfix job-bx wht-shd-bx")

for job in jobs:
  published_date = job.find('span',class_="sim-posted").span.text.replace('\r\n','').replace('  ','')
  
  if 'few'in published_date:
    
    field_name= job.find('h2').text.strip()
    company_name=job.find('h3',class_='joblist-comp-name').text.replace('\r\n','').replace('  ','')
    skills = job.find('span',class_="srp-skills").text.replace('\r\n','').replace('  ','')
    # skills = list(skill)
    more_info=job.header.h2.a['href']
    

    if all([unfamiliar_skill not in skills for unfamiliar_skill in unfamiliar_skills]) and any([familiar_skill in skills for familiar_skill in familiar_skills]):

      
      # with open(f'positions/{index}.txt','w') as f:
        
      print(f'''
      Field: {field_name}
      
      Company Name: {company_name}
      
      Skills Required: {skills}
      
      More_info: {more_info} 
      
      ---------------------------------------------------------------------------------------------------------
      
      ''')
      #  f.write(f' More_info: {more_info} \n')
      # print(f'File saved successfully: {index}')


# if __name__ == '__main__':
#   while True:
#     find_jobs()
#     time_wait = 10
#     print(f'Waiting {time_wait} minutes...')
#     time.sleep(time_wait*60)
   
       