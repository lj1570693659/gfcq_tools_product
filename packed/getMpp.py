import json
import jpype
import warnings
import sys, os
from threading import RLock
 
 
class JarLoader:
    __rlock = RLock()
    __started = False
 
    @classmethod
    def start_jvm(cls):
        if cls.__started:
            return
        with cls.__rlock:
            if not cls.__started:
                jpype.addClassPath("slf4j-api-1.7.31.jar") #the relative path to the jar file
                # 启动JVM
                jpype.startJVM()
                cls.__started = True
 
    @staticmethod
    def load_class(filePath):
        # 加载jar包中的类， 返回类， 注意返回的不是对象。 
        import mpxj
        from java.lang import Double
        from java.text import SimpleDateFormat
        from net.sf.mpxj import ProjectFile, TaskField, Duration, TimeUnit, RelationType
        from net.sf.mpxj.reader import UniversalProjectReader

        obj = UniversalProjectReader()
        project = obj.read(filePath)
        timeForamt = SimpleDateFormat("yyyy-MM-dd HH:mm:ss")
        t = []
        for task in project.getTasks():
            if task.getParentTask() != None:
                child      = task.getChildTasks()
                active = 0
                if task.getActive() == True: 
                    active = 1
                info = {"id": task.getID(),
                        "name": task.getName(), 
                        "active": active,
                        "startTime": task.getStart().toString(), 
                        "endTime": task.getFinish().toString(),
                        "duration": task.getDuration().toString(),
                        "level": task.getOutlineLevel(),
                        "rate": task.getPercentageComplete().toString(),
                        "pid": task.getParentTask().getID(),
                        "cid": len(child), 
                        "pName": task.getParentTask().getName()
                        }                    
                t.append(info)
        return t
        
       
 
    @classmethod
    def close(cls):
        jpype.shutdownJVM()  
        cls.__started = False

if __name__ == "__main__":
    warnings.filterwarnings("ignore")
    os.environ["TF_CPP_MIN_LOG_LEVEL"] = '3'
    if len(sys.argv) >= 3:
        obj = JarLoader()
        try:
            obj.start_jvm()
            content = obj.load_class(sys.argv[1])
            python_string = str(content)
            json_str = json.dumps(python_string,  ensure_ascii=False)
            json_str_2 = json_str.replace("'",'"')
            json_str_2 = json_str_2[1:-1]
            # 打开文件
            file = open(sys.argv[2], "w", encoding='utf-8')
            # 写入内容
            file.write(json_str_2)
            # 关闭文件
            file.close()
            print(json_str_2)
        except BaseException  as e:
            print(e)
            pass
        except ImportError  as err:
            pass
        finally:
            obj.close()
            
